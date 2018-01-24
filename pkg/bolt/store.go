package bolt

import (
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
)

type Store struct {
	// Path where is stored the BoltDB database.
	Path string

	// Services
	UserService    *UserService
	VersionService *VersionService

	db                    *bolt.DB
	checkForDataMigration bool
}

const (
	databaseFileName  = "cloudnativego.db"
	versionBucketName = "version"
	userBucketName    = "users"
)

// NewStore initializes a new Store and the associated services
func NewStore(storePath string) (*Store, error) {
	store := &Store{
		Path:           storePath,
		VersionService: &VersionService{},
		UserService:    &UserService{},
	}
	store.UserService.store = store
	store.VersionService.store = store

	_, err := os.Stat(storePath + "/" + databaseFileName)
	if err != nil && os.IsNotExist(err) {
		store.checkForDataMigration = false
	} else if err != nil {
		return nil, err
	} else {
		store.checkForDataMigration = true
	}

	return store, nil
}

// Open opens and initializes the BoltDB database
func (store *Store) Open() error {
	path := store.Path + "/" + databaseFileName
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	store.db = db
	bucketsToCreate := []string{userBucketName, versionBucketName}
	return db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range bucketsToCreate {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// Close closes the BoltDB database.
func (store *Store) Close() error {
	if store.db != nil {
		return store.db.Close()
	}
	return nil
}

// MigrateData automatically migrate the data based on the DBVersion.
func (store *Store) MigrateData() error {
	if !store.checkForDataMigration {
		err := store.VersionService.StoreDBVersion(cloudnativego.DBVersion)
		if err != nil {
			return err
		}
		return nil
	}
	version, err := store.VersionService.DBVersion()
	if err == cloudnativego.ErrDBVersionNotFound {
		version = 0
	} else if err != nil {
		return err
	}
	if version < cloudnativego.DBVersion {
		log.Printf("Migrating database from version %v to %v.\n", version, cloudnativego.DBVersion)

	}
	return nil
}

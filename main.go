package main

import (
	"fmt"
	"log"

	"github.com/gotoolkit/cloudnativego/pkg/bolt"
	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"
	"github.com/gotoolkit/cloudnativego/pkg/http"
	"github.com/kelseyhightower/envconfig"
)

func initSpec() *cloudnativego.Spec {
	var spec cloudnativego.Spec
	err := envconfig.Process("app", &spec)
	if err != nil {
		log.Fatal(err)
	}
	return &spec

}
func initStore(dataStorePath string) *bolt.Store {
	store, err := bolt.NewStore(dataStorePath)
	if err != nil {
		log.Fatal(err)
	}

	err = store.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = store.MigrateData()
	if err != nil {
		log.Fatal(err)
	}
	return store
}

func main() {
	spec := initSpec()
	store := initStore(spec.DataStorePath)
	defer store.Close()

	var server cloudnativego.Server = &http.Server{
		BindAddress: fmt.Sprintf(":%d", spec.Port),
		UserService: store.UserService,
	}
	log.Printf("Starting CloudNativeGo %s on %s", cloudnativego.APIVersion, fmt.Sprintf(":%d", spec.Port))
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

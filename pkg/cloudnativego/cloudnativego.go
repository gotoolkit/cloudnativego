package cloudnativego

type (
	// Spec represents application env config
	Spec struct {
		Debug         bool
		Port          int
		DataSource    string
		DataStorePath string
	}

	// User represents a user account.
	User struct {
		ID       UserID   `json:"id"`
		Username string   `json:"username"`
		Password string   `json:"password,omitempty"`
		Role     UserRole `json:"role"`
	}
	// UserID represents a user identifier
	UserID int

	// UserRole represents the role of a user
	UserRole int

	// Server defines the interface to serve the API.
	Server interface {
		Start() error
	}

	// UserService represents a service for managing user data.
	UserService interface {
		User(ID UserID) (*User, error)
		UserByUsername(username string) (*User, error)
		Users() ([]User, error)
		UsersByRole(role UserRole) ([]User, error)
		CreateUser(user *User) error
		UpdateUser(ID UserID, user *User) error
		DeleteUser(ID UserID) error
	}

	// TokenData represents the data embedded in a JWT token.
	TokenData struct {
		ID       UserID
		Username string
		Role     UserRole
	}

	// JWTService represents a service for managing JWT tokens.
	JWTService interface {
		GenerateToken(data *TokenData) (string, error)
		ParseAndVerifyToken(token string) (*TokenData, error)
	}
)

const (

	// APIVersion is the version number of the Portainer API.
	APIVersion = "1.0.0"

	// DBVersion is the version number of the cloudnative database.
	DBVersion = 1
)

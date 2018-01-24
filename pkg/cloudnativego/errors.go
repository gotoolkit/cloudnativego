package cloudnativego

// General errors.
const (
	ErrUnauthorized     = Error("Unauthorized")
	ErrResourceNotFound = Error("Unable to find resource")
)

// Version errors.
const (
	ErrDBVersionNotFound = Error("DB version not found")
)

// User errors.
const (
	ErrUserNotFound            = Error("User not found")
	ErrUserAlreadyExists       = Error("User already exists")
	ErrInvalidUsername         = Error("Invalid username. White spaces are not allowed")
	ErrAdminAlreadyInitialized = Error("An administrator user already exists")
	ErrCannotRemoveAdmin       = Error("Cannot remove the default administrator account")
	ErrAdminCannotRemoveSelf   = Error("Cannot remove your own user account. Contact another administrator")
)

// JWT errors.
const (
	ErrSecretGeneration   = Error("Unable to generate secret key")
	ErrInvalidJWTToken    = Error("Invalid JWT token")
	ErrMissingContextData = Error("Unable to find JWT data in request context")
)

// Error represents an application error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}

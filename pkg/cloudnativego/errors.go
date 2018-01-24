package cloudnativego

// General errors.
const (
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

// Error represents an application error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}

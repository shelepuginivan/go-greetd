package greetd

// ErrorType represents type of the error.
type ErrorType string

const (
	// Indicates that authentication failed. This is not a fatal error, and is
	// likely caused by incorrect credentials. Handle as appropriate.
	ErrorAuth ErrorType = "auth_error"

	// A general error. See the error description for more information.
	ErrorGeneral ErrorType = "error"
)

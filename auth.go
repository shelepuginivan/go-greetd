package greetd

// AuthMessageType represents type of the authentication message.
type AuthMessageType string

const (
	// Indicates that input from the user should be visible when they answer this
	// question.
	AuthMessageVisible AuthMessageType = "visible"

	// Indicates that input from the user should be considered secret when they
	// answer this question.
	AuthMessageSecret AuthMessageType = "secret"

	// Indicates that this message is informative, not a question.
	AuthMessageInfo AuthMessageType = "info"

	// Indicates that this message is an error, not a question.
	AuthMessageError AuthMessageType = "error"
)

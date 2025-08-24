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

// ErrorType represents type of the error.
type ErrorType string

const (
	// Indicates that authentication failed. This is not a fatal error, and is
	// likely caused by incorrect credentials. Handle as appropriate.
	ErrorAuth ErrorType = "auth_error"

	// A general error. See the error description for more information.
	ErrorGeneral ErrorType = "error"
)

// ResponseType represents type of the response.
type ResponseType string

const (
	// Indicates that the request succeeded.
	ResponseSuccess ResponseType = "success"

	// Indicates that the request failed.
	ResponseError ResponseType = "error"

	// Indicates that an authentication message needs to be answered to continue
	// through the authentication flow. There are no limits on the number and type
	// of messages that may be required for authentication to succeed, and a
	// greeter should not make any assumptions about the messages. Must be
	// answered with either [RequestPostAuthMessageResponse] or
	// [RequestCancelSession].
	ResponseAuthMessage ResponseType = "auth_message"
)

type Response struct {
	// Type of the response.
	Type ResponseType `json:"type,omitempty"`

	// Type of the occurred error. Present if Type is [ResponseError].
	ErrorType ErrorType `json:"error_type,omitempty"`

	// Description of the error. Present if Type is [ResponseError].
	Description string `json:"description,omitempty"`

	// Type of the auth message. Present if Type is [ResponseAuthMessage].
	AuthMessageType AuthMessageType `json:"auth_message_type,omitempty"`

	// Auth message. Present if Type is [ResponseAuthMessage].
	AuthMessage string `json:"auth_message,omitempty"`
}

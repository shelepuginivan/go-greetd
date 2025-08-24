package greetd

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
	Type ResponseType `json:"type"`

	// Type of the occurred error. Present if Type is [ResponseError].
	ErrorType ErrorType `json:"error_type"`

	// Description of the error. Present if Type is [ResponseError].
	Description string `json:"description"`

	// Type of the auth message. Present if Type is [ResponseAuthMessage].
	AuthMessageType AuthMessageType `json:"auth_message_type"`

	// Auth message. Present if Type is [ResponseAuthMessage].
	AuthMessage string `json:"auth_message"`
}

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

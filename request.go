package greetd

// RequestType represents type of the request.
type RequestType string

const (
	// Creates a session and initiates a login attempted for the given user.
	// The session is ready to be started if a success is returned.
	RequestCreateSession RequestType = "create_session"

	// Answers an authentication message. If the message was informative (info,
	// error), then a response does not need to be set in this message.
	// The session is ready to be started if a success is returned.
	RequestPostAuthMessageResponse RequestType = "post_auth_message_response"

	// Requests for the session to be started using the provided command line,
	// adding the supplied environment to that created by PAM. The session will
	// start after the greeter process terminates.
	RequestStartSession RequestType = "start_session"

	// Cancels the session that is currently under configuration.
	RequestCancelSession RequestType = "cancel_session"
)

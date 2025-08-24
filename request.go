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

// Request represents a request to greetd.
type Request struct {
	// Type of the request.
	Type RequestType `json:"type,omitempty"`

	// Username. Required if Type is [RequestCreateSession].
	Username string `json:"username,omitempty"`

	// Response to the auth message, typically, a password.
	// Required if Type is [RequestPostAuthMessageResponse].
	Response string `json:"response,omitempty"`

	// Command to spawn the session. Required if Type is [RequestStartSession].
	Cmd []string `json:"cmd,omitempty"`

	// Additional environment variables for PAM, in the form of "KEY=VALUE".
	// Required if Type is [RequestStartSession].
	Env []string `json:"env,omitempty"`
}

// NewCreateSessionRequest returns a new [Request] of type
// [RequestCreateSession].
func NewCreateSessionRequest(username string) *Request {
	return &Request{
		Type:     RequestCreateSession,
		Username: username,
	}
}

// NewPostAuthMessageResponseRequest returns a new [Request] of type
// [RequestPostAuthMessageResponse].
func NewPostAuthMessageResponseRequest(response string) *Request {
	return &Request{
		Type:     RequestPostAuthMessageResponse,
		Response: response,
	}
}

// NewStartSessionRequest returns a new [Request] of type
// [RequestStartSession].
func NewStartSessionRequest(cmd []string, env []string) *Request {
	return &Request{
		Type: RequestStartSession,
		Cmd:  cmd,
		Env:  env,
	}
}

// NewCancelSessionRequest returns a new [Request] of type
// [RequestCancelSession].
func NewCancelSessionRequest() *Request {
	return &Request{
		Type: RequestCancelSession,
	}
}

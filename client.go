package greetd

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// Client is a greetd IPC client.
type Client struct {
	conn net.Conn
}

// NewClient connects to the $GREETD_SOCK socket and returns a new instance of
// [Client].
func NewClient() (*Client, error) {
	socket, ok := os.LookupEnv("GREETD_SOCK")
	if !ok {
		return nil, fmt.Errorf("environment variable GREETD_SOCK is not set")
	}

	return NewClientWithSocket(socket)
}

// NewClientWithSocket is like [NewClient], but allows to specify a socket to
// connect to.
func NewClientWithSocket(socket string) (*Client, error) {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

// Close closes the underlying connection.
func (c *Client) Close() error {
	return c.conn.Close()
}

// Read reads raw binary response from greetd socket.
func (c *Client) Read(b []byte) (int, error) {
	return c.conn.Read(b)
}

// Write writes raw binary request to greetd socket.
func (c *Client) Write(b []byte) (int, error) {
	return c.conn.Write(b)
}

// Do sends a request to greetd and returns its response.
//
// An error is returned if client:
//   - fails to send or receive binary message over UNIX socket
//   - is unable to encode request or decode response
//
// If the returned error is nil, response may still contain an error reported
// by greetd, such as authentication error. These errors should be handled
// accordingly.
func (c *Client) Do(request *Request) (*Response, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	length := make([]byte, 4)
	binary.NativeEndian.PutUint32(length, uint32(len(payload)))

	message := make([]byte, len(length)+len(payload))

	copy(message[0:4], length)
	copy(message[4:], payload)

	if _, err := c.conn.Write(message); err != nil {
		return nil, err
	}

	if _, err := c.conn.Read(length); err != nil {
		return nil, err
	}

	responseLength := binary.NativeEndian.Uint32(length)

	buffer := make([]byte, responseLength+1)

	n, err := c.conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.Unmarshal(buffer[0:n], &response); err != nil {
		return nil, err
	}

	return &response, nil
}

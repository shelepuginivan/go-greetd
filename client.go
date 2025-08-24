package greetd

import (
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

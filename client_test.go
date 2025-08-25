package greetd_test

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/shelepuginivan/go-greetd"
)

func TestClient(t *testing.T) {
	t.Run("Basic IPC communication", func(t *testing.T) {
		socket := tmpSocket(t)
		server := createMockSrv(t, socket)
		client, err := NewClientWithSocket(socket)
		assert.NoError(t, err)

		defer func() {
			server.Close()
			client.Close()
			os.Remove(socket)
		}()

		server.requestCount = 2
		server.Accept()

		expectedRequest := NewCreateSessionRequest("user")
		expectedResponse := &Response{Type: ResponseSuccess}
		server.SetExpectations(expectedRequest, expectedResponse)

		actualResponse, err := client.Do(expectedRequest)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, actualResponse)

		expectedRequest = NewPostAuthMessageResponseRequest("strong-password")
		expectedResponse = &Response{
			Type:            ResponseAuthMessage,
			AuthMessageType: AuthMessageError,
			AuthMessage:     "invalid password",
		}
		server.SetExpectations(expectedRequest, expectedResponse)

		actualResponse, err = client.Do(expectedRequest)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, actualResponse)
	})

	t.Run("Socket does not exist", func(t *testing.T) {
		client, err := NewClientWithSocket(tmpSocket(t))
		assert.Nil(t, client)
		assert.Error(t, err)
	})

	t.Run("Using GREETD_SOCK", func(t *testing.T) {
		t.Run("Environment variable is set", func(t *testing.T) {
			socket := tmpSocket(t)
			os.Setenv("GREETD_SOCK", socket)
			server := createMockSrv(t, socket)
			client, err := NewClient()
			assert.NoError(t, err)

			defer func() {
				server.Close()
				client.Close()
				os.Remove(socket)
				os.Unsetenv("GREETD_SOCK")
			}()

			server.requestCount = 1
			server.Accept()

			expectedRequest := NewCreateSessionRequest("user")
			expectedResponse := &Response{Type: ResponseSuccess}
			server.SetExpectations(expectedRequest, expectedResponse)

			actualResponse, err := client.Do(expectedRequest)
			assert.NoError(t, err)
			assert.Equal(t, expectedResponse, actualResponse)
		})

		t.Run("Environment variable is not set", func(t *testing.T) {
			client, err := NewClient()
			assert.Nil(t, client)
			assert.Error(t, err)
		})
	})

	t.Run("Implements ReadWriter", func(t *testing.T) {
		socket := tmpSocket(t)
		server := createMockSrv(t, socket)
		client, err := NewClientWithSocket(socket)
		assert.NoError(t, err)

		defer func() {
			server.Close()
			client.Close()
			os.Remove(socket)
		}()

		server.requestCount = 1
		server.Accept()

		expectedResponse := &Response{Type: ResponseSuccess}
		server.SetExpectations(nil, expectedResponse)
		message := []byte("message")

		n, err := client.Write(message)
		assert.Equal(t, len(message), n)
		assert.NoError(t, err)

		// Payload is as follows
		//
		//   {"type":"success"}
		//
		// The length is 18 + 4 from int32 length prefix.
		buffer := make([]byte, 22)
		n, err = client.Read(buffer)
		assert.Equal(t, n, len(buffer))
		assert.NoError(t, err)
	})
}

func tmpSocket(t *testing.T) string {
	return filepath.Join(t.TempDir(), fmt.Sprintf("go_greetd_test_%d.sock", time.Now().Unix()))
}

// MockGreetdServer is a mock implementation of greetd IPC server used for
// testing.
type MockGreetdServer struct {
	lis  net.Listener
	conn net.Conn

	t                *testing.T
	expectedRequest  *Request
	expectedResponse *Response
	requestCount     int
}

func createMockSrv(t *testing.T, socket string) *MockGreetdServer {
	lis, err := net.Listen("unix", socket)
	if err != nil {
		t.Fatal(err)
	}

	srv := &MockGreetdServer{
		t:   t,
		lis: lis,
	}

	return srv
}

func (srv *MockGreetdServer) SetExpectations(req *Request, res *Response) {
	srv.expectedRequest = req
	srv.expectedResponse = res
}

func (srv *MockGreetdServer) Accept() {
	var err error
	srv.conn, err = srv.lis.Accept()
	assert.NoError(srv.t, err)

	go func() {
		for range srv.requestCount {
			req, err := srv.ReadRequest()

			if srv.expectedRequest != nil {
				assert.NoError(srv.t, err)
				assert.Equal(srv.t, srv.expectedRequest, req)
			}

			err = srv.WriteResponse(srv.expectedResponse)
			assert.NoError(srv.t, err)
		}
	}()
}

func (srv *MockGreetdServer) Close() {
	assert.NoError(srv.t, srv.lis.Close())
}

func (srv *MockGreetdServer) ReadRequest() (*Request, error) {
	lengthBuf := make([]byte, 4)

	if _, err := srv.conn.Read(lengthBuf); err != nil {
		return nil, err
	}

	length := binary.NativeEndian.Uint32(lengthBuf)

	payload := make([]byte, length)

	if _, err := srv.conn.Read(payload); err != nil {
		return nil, err
	}

	var req Request

	if err := json.Unmarshal(payload, &req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (srv *MockGreetdServer) WriteResponse(res *Response) error {
	payload, err := json.Marshal(res)
	if err != nil {
		return err
	}

	length := make([]byte, 4)
	binary.NativeEndian.PutUint32(length, uint32(len(payload)))

	response := make([]byte, len(length)+len(payload))
	copy(response[0:4], length)
	copy(response[4:], payload)

	if _, err := srv.conn.Write(response); err != nil {
		return err
	}
	return nil
}

package greetd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/shelepuginivan/go-greetd"
)

func TestNewCreateSessionRequest(t *testing.T) {
	tests := []struct {
		username string
		want     *Request
	}{
		{
			username: "me",
			want: &Request{
				Type:     RequestCreateSession,
				Username: "me",
			},
		},
		{
			username: "一些文字",
			want: &Request{
				Type:     RequestCreateSession,
				Username: "一些文字",
			},
		},
		{
			username: "",
			want: &Request{
				Type:     RequestCreateSession,
				Username: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := NewCreateSessionRequest(tt.username)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewPostAuthMessageResponseRequest(t *testing.T) {
	tests := []struct {
		response string
		want     *Request
	}{
		{
			response: "password",
			want: &Request{
				Type:     RequestPostAuthMessageResponse,
				Response: "password",
			},
		},
		{
			response: "ютф-8?",
			want: &Request{
				Type:     RequestPostAuthMessageResponse,
				Response: "ютф-8?",
			},
		},
		{
			response: "",
			want: &Request{
				Type:     RequestPostAuthMessageResponse,
				Response: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := NewPostAuthMessageResponseRequest(tt.response)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewStartSessionRequest(t *testing.T) {
	tests := []struct {
		cmd  []string
		env  []string
		want *Request
	}{
		{
			want: &Request{
				Type: RequestStartSession,
			},
		},
		{
			cmd: []string{"/usr/bin/niri-session"},
			want: &Request{
				Type: RequestStartSession,
				Cmd:  []string{"/usr/bin/niri-session"},
			},
		},
		{
			env: []string{"X_GREETD_TEST=some"},
			want: &Request{
				Type: RequestStartSession,
				Env:  []string{"X_GREETD_TEST=some"},
			},
		},
		{
			cmd: []string{"Hyprland"},
			env: []string{"HOME=/home/cutie"},
			want: &Request{
				Type: RequestStartSession,
				Cmd:  []string{"Hyprland"},
				Env:  []string{"HOME=/home/cutie"},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := NewStartSessionRequest(tt.cmd, tt.env)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewCancelSessionRequest(t *testing.T) {
	req := NewCancelSessionRequest()
	assert.Equal(t, &Request{Type: RequestCancelSession}, req)
}

# go-greetd

[![Go Reference](https://pkg.go.dev/badge/github.com/shelepuginivan/go-greetd.svg)](https://pkg.go.dev/github.com/shelepuginivan/go-greetd)
[![Go Report Card](https://goreportcard.com/badge/github.com/shelepuginivan/go-greetd)](https://goreportcard.com/report/github.com/shelepuginivan/go-greetd)
[![License: MIT](https://img.shields.io/badge/License-MIT-00cc00.svg)](https://github.com/shelepuginivan/go-greetd/blob/main/LICENSE)

Package `go-greetd` implements IPC protocol for [greetd](https://git.sr.ht/~kennylevinsen/greetd).
See [`greetd-ipc(7)`](https://man.archlinux.org/man/greetd-ipc.7) for details about the protocol.

## Installation

```sh
go get -u github.com/shelepuginivan/go-greetd
```

## Example usage

```go
package main

import (
	"log"

	"github.com/shelepuginivan/go-greetd"
)

func main() {
	client, err := greetd.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Create new session by providing a username.
	req := greetd.NewCreateSessionRequest("user")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Type == greetd.ResponseError {
		log.Fatal(res.Description)
	}

	// Respond to authentication with a password.
	req = greetd.NewPostAuthMessageResponseRequest("password")
	res, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Type == greetd.ResponseError {
		log.Fatal(res.Description)
	}

	// Start a session.
	req = greetd.NewStartSessionRequest([]string{"niri-session"}, nil)
	res, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Type == greetd.ResponseError {
		log.Fatal(res.Description)
	}
}
```

## Documentation

See package documentation on [pkg.go.dev](https://pkg.go.dev/github.com/shelepuginivan/go-greetd).

## License

[MIT](https://github.com/shelepuginivan/go-greetd/blob/main/LICENSE)

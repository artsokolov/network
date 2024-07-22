package auth

import (
	"bytes"
	"encoding/base64"
	"errors"
	"strings"
)

var (
	ErrInvalidAuthHeader = errors.New("invalid auth header")
)

type Credentials struct {
	Email    string
	Password string
}

func GetCredentials(header string) (*Credentials, error) {
	if strings.TrimSpace(header) == "" {
		return nil, ErrInvalidAuthHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Basic" {
		return nil, ErrInvalidAuthHeader
	}

	decoded, err := base64.StdEncoding.DecodeString(headerParts[1])
	if err != nil {
		return nil, ErrInvalidAuthHeader
	}

	credentials := bytes.Split(decoded, []byte(":"))
	if len(credentials) != 2 {
		return nil, ErrInvalidAuthHeader
	}

	email := string(credentials[0])
	password := string(credentials[1])

	return &Credentials{email, password}, nil
}

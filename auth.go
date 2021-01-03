package main

import (
	"fmt"
	"strings"
)

// Credentials type
type Credentials struct {
	Login     string
	Password  string
	Realm     string
	AccessKey string
	SecretKey string
}

// AuthBackend interface
type AuthBackend interface {
	Authenticate(Credentials) (bool, error)
}

type SimpleAuth struct {
	UserDB map[string]string
}

func (sa SimpleAuth) Authenticate(c Credentials) (bool, error) {
	passwd, ok := sa.UserDB[c.Login]

	if !ok || strings.Compare(passwd, c.Password) != 0 {
		return false, fmt.Errorf("Unable to authenticate: %s", c.Login)
	}

	return true, nil
}

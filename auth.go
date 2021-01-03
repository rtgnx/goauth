package main

import (
	"errors"

	"github.com/msteinert/pam"
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

// PAM Authentication Backend
type PAM struct {
	Service string
}

// Authenticate against pam passwd daabase (Linux Only)
func (p PAM) Authenticate(c Credentials) (bool, error) {

	t, err := pam.StartFunc(p.Service, c.Login, func(s pam.Style, msg string) (string, error) {
		switch s {
		case pam.PromptEchoOff:
			return c.Password, nil
		case pam.PromptEchoOn, pam.ErrorMsg, pam.TextInfo:
			return "", nil
		}
		return "", errors.New("Unrecognized PAM message style")
	})

	if err != nil {
		return err
	}

	if err = t.Authenticate(0); err != nil {
		return false, err
	}

	return true, nil
}

package model

import "errors"

var (
	ErrorMultiplRows = errors.New("err multiple rows found")
	ErrorPassword    = errors.New("Failed error password")
)

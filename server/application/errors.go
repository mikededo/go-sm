package application

import "errors"

var (
	// ErrInvalidRequest Error returned when request received is not valid
	ErrInvalidRequest = errors.New("received request entity is not valid")
)

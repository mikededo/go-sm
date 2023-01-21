package util

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

// Singleton pattern
var lock = &sync.Mutex{}
var validate *validator.Validate

func Validator() *validator.Validate {
	lock.Lock()
	defer lock.Unlock()

	if validate == nil {
		validate = validator.New()
	}

	return validate
}

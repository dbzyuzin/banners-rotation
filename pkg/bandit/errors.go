package bandit

import "errors"

var (
	ErrHandNotFound          = errors.New("hand not found")
	ErrHandsNotSet           = errors.New("hands not set")
	ErrConfirmMoreThenSelect = errors.New("confirm more then select")
)

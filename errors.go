package gortifacts

import (
	"errors"
)

var (
	ErrorInvalidInputData = errors.New("invalid input data")

	ErrorCharacterIsLocked      = errors.New("character is locked")
	ErrorCharacterAtDestination = errors.New("character already at destination")
	ErrorCharacterNotFound      = errors.New("character not found")
	ErrorCharacterCooldown      = errors.New("character in cooldown")
)

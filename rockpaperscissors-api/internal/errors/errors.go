package errors

import (
	_errors "errors"
)

var (
	ErrInvalidGameID       = _errors.New("invalidGameID")
	ErrGameNotFound        = _errors.New("gameNotFound")
	ErrPlayerNotFound      = _errors.New("playerNotFound")
	ErrMaxPlayersReached   = _errors.New("maxPlayersReached")
	ErrPlayerAlreadyJoined = _errors.New("playerAlreadyJoined")
	ErrMissingGame         = _errors.New("missingGame")
	ErrPlayerNotJoined     = _errors.New("playerNotJoined")
	ErrInvalidCommand      = _errors.New("invalidCommand")
)

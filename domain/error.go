package domain

import "errors"

var (
	ErrNoSuchEntity    = errors.New("no such entity")
	ErrInvalidArgument = errors.New("invalid argument")
)

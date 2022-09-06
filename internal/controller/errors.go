package controller

import "errors"

var (
	ErrPathValue      = errors.New("Not exist [/memberships/:id] id in path")
	ErrInvalidRequest = errors.New("err Invalid Input")
)

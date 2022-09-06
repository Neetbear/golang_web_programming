package repository

import "errors"

var (
	ErrDuplicateName = errors.New("err duplicate username")
	ErrNotExist      = errors.New("err not exist MemberShip")
)

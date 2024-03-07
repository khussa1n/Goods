package custom_error

import "errors"

var (
	ErrProjectNotFound  = errors.New("errors.project.notfound")
	ErrGoodNotFound     = errors.New("errors.good.notFound")
	ErrInvalidInputBody = errors.New("errors.invalidInputBody")
	ErrInvalidURLHead   = errors.New("errors.invalidURLHead")
)

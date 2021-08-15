package validation

import (
	"errors"
)

var (
	errInvalidValidatorInstance = errors.New("invalid validator instancce")
)

const (
	InstanceGoPlayground int = iota
)

func NewValidatorFactory(instance int) (Validator, error) {
	switch instance {
	case InstanceGoPlayground:
		return NewGoPlayground()
	default:
		return nil, errInvalidValidatorInstance
	}
}

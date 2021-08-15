package database

import (
	"errors"
)

const (
	InstanceGorm int = iota
)

var (
	errInvalidORMDatabaseInstance = errors.New("invalid gorm db instance")
)

func NewDatabaseORMFactory(instance int) (Gorm, error) {
	switch instance {
	case InstanceGorm:
		return NewGormHandler(newConfigPostgres())
	default:
		return nil, errInvalidORMDatabaseInstance
	}
}

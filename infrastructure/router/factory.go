package router

import (
	"auth-service-clean2/infrastructure/database"
	"auth-service-clean2/infrastructure/log"
	"auth-service-clean2/infrastructure/validation"
	"errors"
	"time"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceEcho int = iota
)

func NewWebServerFactoryGorm(
	instance int,
	log log.Logger,
	dbGorm database.Gorm,
	validator validation.Validator,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	case InstanceEcho:
		return newEchoServerGorm(dbGorm, validator, port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}

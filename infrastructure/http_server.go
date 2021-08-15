package infrastructure

import (
	"auth-service-clean2/infrastructure/database"
	"auth-service-clean2/infrastructure/log"
	"auth-service-clean2/infrastructure/router"
	"auth-service-clean2/infrastructure/validation"
	"strconv"
	"time"
)

type config struct {
	appName       string
	logger        log.Logger
	validator     validation.Validator
	dbGorm        database.Gorm
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) Name(name string) *config {
	c.appName = name
	return c
}

func (c *config) Logger(instance int) *config {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Succesfully configured log")
	return c
}

func (c *config) DbGorm(instance int) *config {
	db, err := database.NewDatabaseORMFactory(instance)
	if err != nil {
		c.logger.Fatalln(err, "could not make a connection to the daatabase")
	}

	c.logger.Infof("successfully connected to the SQL database")
	c.dbGorm = db
	return c
}

func (c *config) Validator(instance int) *config {
	v, err := validation.NewValidatorFactory(instance)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured validator")

	c.validator = v
	return c
}

func (c *config) WebServerGorm(instance int) *config {
	s, err := router.NewWebServerFactoryGorm(
		instance,
		c.logger,
		c.dbGorm,
		c.validator,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("successfully configued router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}

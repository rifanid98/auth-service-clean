package main

import (
	"auth-service-clean2/infrastructure"
	"auth-service-clean2/infrastructure/database"
	"auth-service-clean2/infrastructure/log"
	"auth-service-clean2/infrastructure/router"
	"auth-service-clean2/infrastructure/validation"
	"os"
	"time"
)

func main() {
	app := infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceZapLogger).
		Validator(validation.InstanceGoPlayground).
		DbGorm(database.InstanceGorm)

	app.WebServerPort("8080").
		WebServerGorm(router.InstanceEcho).
		Start()
}

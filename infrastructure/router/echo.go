package router

import (
	"auth-service-clean2/adapter/handler"
	"auth-service-clean2/adapter/presenter"
	"auth-service-clean2/domain"
	"auth-service-clean2/infrastructure/database"
	"auth-service-clean2/infrastructure/persistence"
	"auth-service-clean2/infrastructure/validation"
	"auth-service-clean2/usecase"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

type echoEngineGorm struct {
	router     *echo.Echo
	db         database.Gorm
	validator  validation.Validator
	port       Port
	ctxTimeout time.Duration
}

func newEchoServerGorm(
	db database.Gorm,
	validator validation.Validator,
	port Port,
	t time.Duration,
) *echoEngineGorm {
	return &echoEngineGorm{
		router:     echo.New(),
		db:         db,
		validator:  validator,
		port:       port,
		ctxTimeout: t,
	}
}

func (e *echoEngineGorm) Listen() {
	e.setAppHandlers()

	e.db.AutoMigrate(&domain.Patient{})
	// e.db.AutoMigrate(&domain.Patient{}, &domain.Doctor{}, &domain.Partner{})

	e.router.Logger.Fatal(e.router.Start(fmt.Sprintf(":%v", e.port)))
}

func (e *echoEngineGorm) setAppHandlers() {
	e.router.POST("/auth/register", e.buildRegisterHandler)
}

func (e *echoEngineGorm) buildRegisterHandler(c echo.Context) error {
	var (
		uc = usecase.NewRegisterPatientInteractor(
			persistence.NewPatientGorm(e.db),
			presenter.NewRegisterPatientPresenter(),
			e.ctxTimeout,
		)

		act = handler.NewRegisterPatientHandler(uc, e.router.Logger, e.validator)
	)

	return act.Execute(c)
}

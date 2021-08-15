package handler

import (
	"auth-service-clean2/adapter/response"
	"auth-service-clean2/infrastructure/validation"
	"auth-service-clean2/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterPatientHandler struct {
	log       echo.Logger
	uc        usecase.RegisterPatientUseCase
	validator validation.Validator

	logKey, logMsg string
}

func NewRegisterPatientHandler(uc usecase.RegisterPatientUseCase, log echo.Logger, v validation.Validator) RegisterPatientHandler {
	return RegisterPatientHandler{
		log:       log,
		uc:        uc,
		validator: v,
		logKey:    "health_check",
		logMsg:    "health checking",
	}
}

func (h RegisterPatientHandler) Execute(c echo.Context) error {
	var input usecase.RegisterPatientInput
	err := c.Bind(&input)
	if err != nil {
		h.log.Errorf("failed to bind payload", err)
		return c.JSON(http.StatusBadRequest, response.NewError(response.ErrInvalidInput, http.StatusBadRequest))
	}

	if errs := h.ValidateInput(input); len(errs) > 0 {
		h.log.Info("invalid payload")
		return c.JSON(http.StatusBadRequest, response.NewErrorMessage(errs, http.StatusBadRequest))
	}

	output, err := h.uc.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError(response.ErrInternalServerError, http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, response.NewSuccess(output, http.StatusOK))
}

func (h RegisterPatientHandler) ValidateInput(input usecase.RegisterPatientInput) []string {
	var msgs []string

	// for k, v := range enum.Patient {
	// 	if pass := h.validator.ValidateEnumString(v, h.validator.GetFieldValue(input, k)); !pass {
	// 		// strings.ToUpper(string(k[0]))+k[1:]
	// 		// word -> Word
	// 		fieldName := strings.ToUpper(string(k[0])) + k[1:]
	// 		msgs = append(msgs, fmt.Sprintf("%s must be one of %v", fieldName, v))
	// 	}
	// }

	err := h.validator.Validate(input)
	if err != nil {
		msgs = append(msgs, h.validator.Messages()...)
	}

	return msgs
}

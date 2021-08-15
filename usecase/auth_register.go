package usecase

import (
	"auth-service-clean2/common/enum"
	"auth-service-clean2/domain"
	"time"
)

type (
	RegisterPatientUseCase interface {
		Execute(RegisterPatientInput) (RegisterPatientOutput, error)
	}

	RegisterPatientPresenter interface {
		Output(*domain.Patient) RegisterPatientOutput
	}

	RegisterPatientInput struct {
		Type  enum.UserType `json:"type" validate:"required"`
		Name  string        `json:"name" validate:"required,min=3"`
		Email string        `json:"email" validate:"required,email"`
		Phone string        `json:"phone" validate:"required,min=10"`
	}

	RegisterPatientOutput struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	registerInteractor struct {
		repo       domain.PatientRepository
		presenter  RegisterPatientPresenter
		ctxTimeout time.Duration
	}
)

func NewRegisterPatientInteractor(
	repo domain.PatientRepository,
	presenter RegisterPatientPresenter,
	t time.Duration,
) RegisterPatientUseCase {
	return registerInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (i registerInteractor) Execute(input RegisterPatientInput) (RegisterPatientOutput, error) {
	patient := domain.NewPatient()
	patient.Name = &input.Name
	patient.Email = &input.Email
	patient.Phone = &input.Phone
	patient.IsParent = true
	patient.IsVisible = true

	registered, err := i.repo.Create(patient)

	if err != nil {
		return RegisterPatientOutput{}, err
	}

	return i.presenter.Output(registered), nil
}

package presenter

import (
	"auth-service-clean2/domain"
	"auth-service-clean2/usecase"
)

type registerPatientPresenter struct{}

func NewRegisterPatientPresenter() usecase.RegisterPatientPresenter {
	return registerPatientPresenter{}
}

func (c registerPatientPresenter) Output(patient *domain.Patient) usecase.RegisterPatientOutput {
	return usecase.RegisterPatientOutput{
		Name:  *patient.Name,
		Email: *patient.Email,
		Phone: *patient.Phone,
	}
}

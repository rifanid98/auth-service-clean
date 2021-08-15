package persistence

import (
	"auth-service-clean2/domain"
	"auth-service-clean2/infrastructure/database"
)

type PatientGorm struct {
	db database.Gorm
}

func NewPatientGorm(db database.Gorm) PatientGorm {
	return PatientGorm{db: db}
}

func (p PatientGorm) Create(patient *domain.Patient) (*domain.Patient, error) {
	result, err := p.db.Create(patient, "patients")
	if err != nil {
		return domain.NewPatient(), err
	}

	return result.(*domain.Patient), nil
}

func (p PatientGorm) GetAll() ([]domain.Patient, error) {
	return nil, nil
}

func (p PatientGorm) FindByPatientId(pid int64) (*domain.Patient, error) {
	return domain.NewPatient(), nil
}

func (p PatientGorm) FindByEmailAndPhone(email, phone string) (*domain.Patient, error) {
	return domain.NewPatient(), nil
}

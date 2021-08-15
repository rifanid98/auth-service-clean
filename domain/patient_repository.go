package domain

type PatientRepository interface {
	GetAll() ([]Patient, error)
	FindByPatientId(pid int64) (*Patient, error)
	FindByEmailAndPhone(email, phone string) (*Patient, error)
	Create(patient *Patient) (*Patient, error)
}

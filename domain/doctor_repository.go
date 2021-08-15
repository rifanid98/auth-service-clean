package domain

type DoctorRepository interface {
	GetAll() ([]Doctor, error)
	FindByDoctorId(pid int64) (Doctor, error)
	FindByEmailAndPhone(email, phone string) (Doctor, error)
	Create(patient *Doctor) (Doctor, error)
}

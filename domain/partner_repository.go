package domain

type PartnerRepository interface {
	GetAll() ([]Partner, error)
	FindByPartnerId(pid int64) (Partner, error)
	FindByEmailAndPhone(email, phone string) (Partner, error)
	Create(patient *Partner) (Partner, error)
}

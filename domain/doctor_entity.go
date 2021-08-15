package domain

import (
	"auth-service-clean2/common/enum"
	"auth-service-clean2/common/types"
	"time"
)

type Doctor struct {
	Id                        int64                   `gorm:"primary_key;auto_increment" json:"id,omitempty" validate:"min=1"`
	Name                      string                  `json:"name,omitempty" validate:"min=3,max=300"`
	Gender                    enum.Gender             `json:"gender,omitempty" sql:"type:gender"`
	BirthDate                 time.Time               `json:"birth_date,omitempty"`
	BirthPlace                string                  `json:"birth_place,omitempty"`
	Religion                  enum.Religion           `json:"religion,omitempty" sql:"type:religion"`
	ReligionDetail            string                  `json:"religion_detail,omitempty" validate:"min=3"`
	Phone                     string                  `json:"phone,omitempty" validate:"min=11"`
	Email                     string                  `json:"email,omitempty" validate:"email"`
	Otp                       string                  `json:"otp,omitempty" validate:"min=4,max=4"`
	OtpCreated                *time.Time              `json:"otp_created,omitempty"`
	Status                    enum.Status             `json:"status,omitempty" sql:"type:status"`
	IsVisible                 bool                    `json:"is_visible,omitempty"`
	FcmToken                  string                  `json:"fcm_token,omitempty"`
	Device                    enum.Device             `json:"device,omitempty" sql:"type:device"`
	DeviceInfo                string                  `json:"device_info,omitempty" validate:"min=3"`
	PriceLevelId              int64                   `json:"price_level_id,omitempty" validate:"min=1"`
	University                string                  `json:"university,omitempty" validate:"min=3,max=300"`
	GraduationYear            *time.Time              `json:"graduation_year,omitempty"`
	SIP                       string                  `json:"sip,omitempty" validate:"min=3,max=300"`
	Experience                string                  `json:"experience,omitempty" validate:"min=3,max=300"`
	Clinic                    string                  `json:"clinic,omitempty" validate:"min=3,max=300"`
	IsOnService               bool                    `json:"is_on_service,omitempty"`
	PartnerId                 int64                   `json:"partner_id,omitempty" validate:"min=1"`
	DoctorType                string                  `json:"doctor_type,omitempty" validate:"min=3,max=300"`
	Description               string                  `json:"description,omitempty" validate:"min=3,max=300"`
	IconId                    string                  `json:"icon_id,omitempty" validate:"min=3,max=300"`
	IconUrl                   string                  `json:"icon_url,omitempty" validate:"min=3,max=300"`
	PictureUrl                string                  `json:"picture_url,omitempty" sql:"type:point" validate:"min=3,max=300"`
	Location                  types.Point             `json:"location,omitempty" validate:""`
	Address                   string                  `json:"address,omitempty" validate:"min=3,max=300"`
	RegistrationNumber        int64                   `json:"registration_number,omitempty" validate:"min=1"`
	RegistrationCode          string                  `json:"registration_code,omitempty" validate:"min=3,max=300"`
	RegistrationInterviewDate *time.Time              `json:"registration_interview_date,omitempty"`
	RegistrationExpire        *time.Time              `json:"registration_expire,omitempty"`
	RegistrationStatus        enum.RegistrationStatus `json:"registration_status,omitempty" sql:"type:registration_status"`
	VerivicationToken         string                  `json:"verification_token,omitempty" validate:"min=3,max=300"`
	CreatedAt                 *time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt                 *time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

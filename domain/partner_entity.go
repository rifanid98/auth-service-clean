package domain

import (
	"auth-service-clean2/common/enum"
	"auth-service-clean2/common/types"
	"time"
)

type Partner struct {
	Id                        int64                   `gorm:"primary_key;auto_increment" json:"id,omitempty" validate:"min=1"`
	Name                      string                  `json:"name,omitempty" validate:"min=3,max=300"`
	Phone                     string                  `json:"phone,omitempty" validate:"min=11"`
	Email                     string                  `json:"email,omitempty" validate:"email"`
	Otp                       string                  `json:"otp,omitempty" validate:"min=4,max=4"`
	OtpCreated                *time.Time              `json:"otp_created,omitempty"`
	Status                    enum.Status             `json:"status,omitempty" sql:"type:status"`
	IsVisible                 bool                    `json:"is_visible,omitempty"`
	FcmToken                  string                  `json:"fcm_token,omitempty"`
	Device                    enum.Device             `json:"device,omitempty" sql:"type:device"`
	DeviceInfo                string                  `json:"device_info,omitempty" validate:"min=3"`
	PartnerType               string                  `json:"partner_type,omitempty" validate:"min=3,max=300"`
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

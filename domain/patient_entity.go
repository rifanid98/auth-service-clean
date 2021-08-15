package domain

import (
	"auth-service-clean2/common/enum"
	"time"
)

type Patient struct {
	Id                  int64          `gorm:"primary_key;auto_increment" json:"id,omitempty" validate:"min=1"`
	Name                *string        `json:"name,omitempty" validate:"min=3,max=300"`
	Gender              *enum.Gender   `json:"gender,omitempty" sql:"type:gender"`
	BirthDate           time.Time      `json:"birth_date,omitempty"`
	BirthPlace          *string        `json:"birth_place,omitempty"`
	Religion            *enum.Religion `json:"religion,omitempty" sql:"type:religion"`
	ReligionDetail      *string        `json:"religion_detail,omitempty" validate:"min=3"`
	Phone               *string        `json:"phone,omitempty" validate:"min=11"`
	Email               *string        `json:"email,omitempty" validate:"email"`
	Profession          *string        `json:"profession,omitempty" validate:"min=3"`
	AllergyHistory      *string        `json:"allergy_history,omitempty"`
	PrivateDoctor       *string        `json:"private_doctor,omitempty" validate:"min=3"`
	PrivateDoctorPhone  *string        `json:"private_doctor_phone,omitempty" validate:"min=11"`
	HealthInsurance     *string        `json:"health_insurance,omitempty" validate:"min=3"`
	MedicalRecordNumber *string        `json:"medical_record_number,omitempty" validate:"min=7"`
	OTP                 *string        `json:"otp,omitempty" validate:"min=4,max=4"`
	OTPCreated          time.Time      `json:"otp_created,omitempty"`
	Status              *enum.Status   `json:"status,omitempty" sql:"type:status"`
	IsParent            bool           `json:"is_parent,omitempty"`
	IsVisible           bool           `json:"is_visible,omitempty"`
	FcmToken            *string        `json:"fcm_token,omitempty"`
	Device              *enum.Device   `json:"device,omitempty" sql:"type:device"`
	DeviceInfo          *string        `json:"device_info,omitempty" validate:"min=3"`
	VerificationToken   *string        `json:"verification_token,omitempty" validate:"min=3"`
	VerificationExpired time.Time      `json:"verification_expired,omitempty"`
	PictureUrl          *string        `json:"picture_url,omitempty"`
	FacebookId          *string        `json:"facebook_id,omitempty"`
	GoogleId            *string        `json:"google_id,omitempty"`
	CreatedAt           time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt           time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func NewPatient() *Patient {
	return &Patient{}
}

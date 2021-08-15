package domain

import "auth-service-clean2/common/enum"

type Register struct {
	Type  enum.UserType `json:"type" validate:"required"`
	Name  string        `json:"name" validate:"required,min=3"`
	Email string        `json:"email" validate:"required,email"`
	Phone string        `json:"phone" validate:"required,min=10"`
}

type RequestOTP struct {
	Type  enum.UserType `json:"type" validate:"required"`
	Phone string        `json:"phone" validate:"required,min=10,max=15"`
}

type ValidateOTP struct {
	Type       enum.UserType `json:"type" validate:"required"`
	OTP        string        `json:"otp" validate:"required,min=4,max=4"`
	Phone      string        `json:"phone" validate:"required,min=10,max=15"`
	Device     enum.Device   `json:"device" validate:"required"`
	DeviceInfo string        `json:"device_info" validate:"required"`
}

package enum

import "database/sql/driver"

var (
	religions, statuses, device, gender, userTypes, registrationStatus []interface{}
)

func init() {
	setReligionsEnum()
	setStatusesEnum()
	setDeviceEnum()
	setGenderEnum()
	setUserTypesEnum()
	setRegistrationStatusEnum()
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (g *Gender) Scan(value interface{}) error {
	*g = Gender(value.([]byte))
	return nil
}

func (p Gender) Value() (driver.Value, error) {
	return string(p), nil
}

func (Gender) TableName() string {
	return "gender"
}

func setGenderEnum() {
	gender = make([]interface{}, 0)
	gender = append(gender, "female")
	gender = append(gender, "male")
	gender = append(gender, "other")
}

type Status string

const (
	StatusPending  Status = "pending"
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (g *Status) Scan(value interface{}) error {
	*g = Status(value.([]byte))
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}

func (Status) TableName() string {
	return "status"
}

func setStatusesEnum() {
	statuses = make([]interface{}, 0)
	statuses = append(statuses, "active")
	statuses = append(statuses, "inactive")
	statuses = append(statuses, "pending")
}

type Device string

const (
	IOS     Device = "ios"
	Android Device = "android"
	Web     Device = "web"
)

func (g *Device) Scan(value interface{}) error {
	*g = Device(value.([]byte))
	return nil
}

func (p Device) Value() (driver.Value, error) {
	return string(p), nil
}

func (Device) TableName() string {
	return "device"
}

func setDeviceEnum() {
	device = make([]interface{}, 0)
	device = append(device, "android")
	device = append(device, "ios")
	device = append(device, "web")
}

type Religion string

const (
	Islam        Religion = "islam"
	Protestant   Religion = "protestant"
	Catholic     Religion = "catholic"
	Hindu        Religion = "hindu"
	Buddha       Religion = "buddha"
	Confucianism Religion = "confucianism"
	Other        Religion = "other"
)

func (g *Religion) Scan(value interface{}) error {
	*g = Religion(value.([]byte))
	return nil
}

func (p Religion) Value() (driver.Value, error) {
	return string(p), nil
}

func (Religion) TableName() string {
	return "religion"
}

func setReligionsEnum() {
	religions = make([]interface{}, 0)
	religions = append(religions, "buddha")
	religions = append(religions, "catholic")
	religions = append(religions, "confucianism")
	religions = append(religions, "hindu")
	religions = append(religions, "islam")
	religions = append(religions, "other")
}

type UserType string

const (
	DoctorUser  UserType = "doctor"
	PatientUser UserType = "patient"
	PartnerUser UserType = "partner"
)

func (g *UserType) Scan(value interface{}) error {
	*g = UserType(value.([]byte))
	return nil
}

func (p UserType) Value() (driver.Value, error) {
	return string(p), nil
}

func (UserType) TableName() string {
	return "type"
}

func setUserTypesEnum() {
	userTypes = make([]interface{}, 0)
	userTypes = append(userTypes, "doctor")
	userTypes = append(userTypes, "partner")
	userTypes = append(userTypes, "patient")
}

type RegistrationStatus string

const (
	RegAccepted RegistrationStatus = "doctor"
	RegRejected RegistrationStatus = "doctor"
	RegPending  RegistrationStatus = "pending"
)

func (g *RegistrationStatus) Scan(value interface{}) error {
	*g = RegistrationStatus(value.([]byte))
	return nil
}

func (p RegistrationStatus) Value() (driver.Value, error) {
	return string(p), nil
}

func (RegistrationStatus) TableName() string {
	return "type"
}

func setRegistrationStatusEnum() {
	registrationStatus = make([]interface{}, 0)
	registrationStatus = append(registrationStatus, "doctor")
	registrationStatus = append(registrationStatus, "partner")
	registrationStatus = append(registrationStatus, "patient")
}

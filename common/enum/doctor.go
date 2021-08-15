package enum

var (
	Doctor map[string][]interface{}
)

func init() {
	setDoctorEnum()
}

func setDoctorEnum() {
	Doctor = map[string][]interface{}{
		"Gender":             gender,
		"Religions":          religions,
		"Statuses":           statuses,
		"Device":             device,
		"RegistrationStatus": registrationStatus,
	}
}

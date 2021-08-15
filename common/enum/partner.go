package enum

var (
	Partner map[string][]interface{}
)

func init() {
	setPartnerEnum()
}

func setPartnerEnum() {
	Partner = map[string][]interface{}{
		"Statuses":           statuses,
		"Device":             device,
		"RegistrationStatus": registrationStatus,
	}
}

package enum

var (
	Patient map[string][]interface{}
)

func init() {
	setPatientEnum()
}

func setPatientEnum() {
	Patient = map[string][]interface{}{
		"Gender":    gender,
		"Religions": religions,
		"Statuses":  statuses,
		"Device":    device,
	}
}

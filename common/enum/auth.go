package enum

var (
	Register map[string][]interface{}
)

func init() {
	setAuthEnum()
}

func setAuthEnum() {
	Register = map[string][]interface{}{
		"Type": userTypes,
	}
}

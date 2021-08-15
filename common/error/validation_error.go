package error

type RequiredFieldError struct {
	message string
}

func NewRequiredFieldError(msg string) *RequiredFieldError {
	return &RequiredFieldError{msg}
}

func (e *RequiredFieldError) Error() string {
	return e.message
}

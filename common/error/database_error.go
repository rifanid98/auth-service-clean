package error

type FailedSaveError struct {
	message string
}

func NewFailedSaveError(msg string) *FailedSaveError {
	return &FailedSaveError{msg}
}

func (e *FailedSaveError) Error() string {
	return e.message
}

type NoDataError struct {
	message string
}

func NewNoDataError(msg string) *NoDataError {
	return &NoDataError{msg}
}

func (e *NoDataError) Error() string {
	return e.message
}

type DuplicateDataError struct {
	message string
}

func NewDuplicateDataError(msg string) *DuplicateDataError {
	return &DuplicateDataError{msg}
}

func (e *DuplicateDataError) Error() string {
	return e.message
}

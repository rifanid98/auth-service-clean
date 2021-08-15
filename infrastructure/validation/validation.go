package validation

import "auth-service-clean2/common/types"

type Validator interface {
	Validate(payload interface{}) error
	ValidatePartial(payload interface{}, fields ...string) error
	ValidatePointValue(point types.Point) error
	ValidateEnumString(enum []interface{}, value interface{}) bool
	GetFieldValue(v interface{}, field string) string
	Messages() []string
}

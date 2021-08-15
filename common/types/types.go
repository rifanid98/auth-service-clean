package types

type Point struct {
	X string `json:"x" validate:"numeric"`
	Y string `json:"y" validate:"numeric"`
}

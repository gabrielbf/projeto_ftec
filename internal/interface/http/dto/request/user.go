package request

type CreateUser struct {
	FirstName string `json:"first_name" validate:"required" example:"Lucas"`     //
	LastName  string `json:"last_name" validate:"required" example:"Montenegro"` //
}

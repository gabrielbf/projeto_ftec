package request

type CreateAccount struct {
	Email    string `json:"email" validate:"required,email" example:"user@email.com"` //
	Password string `json:"password" validate:"required" example:"password"`          //
	Type     string `json:"type" validate:"required" example:"USER"`                  //
}

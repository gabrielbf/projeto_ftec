package request

type CreateLocation struct {
	Country string `json:"country" validate:"required" example:"Brasil"`//
	State string `json:"state" validate:"required" example:"RS"` //
	City string `json:"city" validate:"required" example:"Porto Alegre"` //
	Neighborhood string `json:"neighborhood" validate:"required" example:"Centro"` //
	Street string `json:"city" validate:"required" example:"Independência"` //
	Number string `json:"city" validate:"required" example:"138"` //
	Complement string `json:"city" validate:"required" example:"Apartamento 1001"` //
	Cep string `json:"city" validate:"required" example:"90994-90"` //
	Coordinates string `json:"city" validate:"required" example:"-29.965017824167994, -51.32133750726387"` //
	Description string `json:"city" validate:"required" example:"Descrição"` //

}
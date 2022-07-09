package request

type CreateAccount struct {
	Email        string `json:"email" validate:"required,email" example:"user@email.com"` //
	Password     string `json:"password" validate:"required" example:"password"`          //
	Type         string `json:"type" validate:"required" example:"USER"`                  //
	Name         string `json:"name" validate:"required" example:"Lucas"`                 //
	PhoneNumber  string `json:"phone_number" validate:"required" example:"5551988888888"` //
	Country      string `json:"country" validate:"required" example:"BRASIL"`             //
	State        string `json:"state" validate:"required" example:"Rio Grande do Sul"`    //
	City         string `json:"city" validate:"required" example:"Porto Alegre"`          //
	Neighborhood string `json:"neighborhood" validate:"required" example:"Centro"`        //
	Street       string `json:"street" validate:"required" example:"Salgado Filho"`       //
	Number       string `json:"number" validate:"required" example:"2022"`                //
}

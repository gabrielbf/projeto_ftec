package request

type CreatePartner struct {
	Name string `json:"name" validate:"required" example:"Petshop do Lucas"` //
}

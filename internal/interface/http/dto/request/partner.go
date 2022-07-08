package request

type CreatePartner struct {
	ReferenceUUID string `json:"reference_uuid" validate:"required,uuid" example:"6fe2d9d5-687d-47da-aaf2-a31d88aa70ca"` // UUID sent by client and used as a client generated id. It will be the external reference for the requested partner.
}

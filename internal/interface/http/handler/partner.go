package handler

import (
	"net/http"

	"github.com/ftec-project/internal/domain/partner"
	"github.com/ftec-project/internal/interface/http/dto/request"
	"github.com/ftec-project/internal/interface/http/dto/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"schneider.vip/problem"
)

func MakePartnerHandler(api *echo.Echo, partnerService partner.Service) {
	api.POST("/v1/partners", CreatePartner(partnerService))
}

// Partner godoc
// @Router /partners [POST]
// @Summary Request an Partner creation
// @Description Partner Description.
// @Tags Partner
// @Accept json
// @Param CreatePartner body request.CreatePartner{} true "Requested body for Partner Creation"
// @Produce json
// @Success 201 {object} response.Partner{} "Partner request successfully created"
// @Header 201 {string} Location "The endpoint location for the created ressource. E.g /partners/da74025d-07cf-4347-8e97-3073e83adc82"
// @Failure 400  {object}  response.Error
// @Failure 500  {object}  response.Error
func CreatePartner(partnerService partner.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("creating Partner")

		var createPartnerRequest request.CreatePartner
		var createdPartnerResponse response.Partner

		if err := c.Bind(&createPartnerRequest); err != nil {
			log.Info().Interface("createPartnerRequest", createPartnerRequest).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		createDTO := partner.CreateDTO{
			Name: createPartnerRequest.Name,
		}

		createdPartner, err := partnerService.CreatePartner(c.Request().Context(), createDTO)
		if err != nil {
			log.Error().Interface("createPartnerRequest", createPartnerRequest).Msg("error on Partner creation")
			status := http.StatusInternalServerError
			prob := problem.New(
				problem.Title("UNEXPECTED_ERROR"),
				problem.Detail("Internal Server Error"),
				problem.Status(status),
			)
			return c.JSON(status, prob)
		}

		log.Info().Msg("retrieving created partner")

		createdPartnerResponse.FromPartner(createdPartner)
		return c.JSON(http.StatusCreated, createdPartnerResponse)
	}
}

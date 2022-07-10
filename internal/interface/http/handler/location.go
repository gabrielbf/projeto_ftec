package handler

import (
	"net/http"

	"github.com/ftec-project/internal/domain/location"
	"github.com/ftec-project/internal/interface/http/dto/request"
	"github.com/ftec-project/internal/interface/http/dto/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"schneider.vip/problem"
)

func MakeLocationHandler(api *echo.Echo, locationService location.Service){
	api.POST("/v1/location", CreateLocation(locationService))
}

func CreateLocation(locationService location.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("creating location")

		var createLocationRequest request.CreateLocation
		var createdLocationResponse response.Location

		if err := c.Bind(&createLocationRequest); err != nil {
			log.Info().Interface("createLocationRequest", createLocationRequest).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		createDTO := location.CreateDTO{
			Country: createLocationRequest.Country,
			State: createLocationRequest.State,
			City: createLocationRequest.City,
			Neighborhood: createLocationRequest.Neighborhood,
			Street: createLocationRequest.Street,
			Number: createLocationRequest.Number,
			Complement: createLocationRequest.Complement,
			Cep: createLocationRequest.Cep,
			Coordinates: createLocationRequest.Coordinates,
			Description: createLocationRequest.Description,
		}

		createdLocation, err := locationService.CreateLocation(c.Request().Context(), createDTO)
		if err != nil {
			log.Error().Interface("createLocationRequest", createLocationRequest).Msg("error on Location creation")
			status := http.StatusInternalServerError
			prob := problem.New(
				problem.Title("UNEXPECTED_ERROR"),
				problem.Detail("Internal Server Error"),
				problem.Status(status),
			)
			return c.JSON(status, prob)
		}

		log.Info().Msg("retrieving created location")

		createdLocationResponse.FromLocation(createdLocation)
		return c.JSON(http.StatusCreated, createdLocationResponse)
	}
}
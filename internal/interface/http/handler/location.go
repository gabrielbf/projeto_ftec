package handler

import (
	"net/http"

	"github.com/ftec-project/internal/domain/location"
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
		var createdLocationResponse request.Location

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
			Country: createAccountRequest.Country
			State: createAccountRequest.State
			City: createAccountRequest.City
			Neighborhood: createAccountRequest.Neighborhood
			Street: createAccountRequest.Street
			Number: createAccountRequest.Number
			Complement: createAccountRequest.Complement
			Cep: createAccountRequest.Cep
			Coordinates: createAccountRequest.Coordinates
			Description: createAccountRequest.Description
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

		createLocationResponse.FromAccount(createdLocation)
		return c.JSON(http.StatusCreated, createdLocationResponse)
	}
}
package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/sample-project/internal/domain/sample"
	"github.com/sample-project/internal/infra/constants"
	"github.com/sample-project/internal/interface/http/dto/request"
	"github.com/sample-project/internal/interface/http/dto/response"
	"gorm.io/gorm"
	"schneider.vip/problem"
)

func MakeSampleHandler(api *echo.Echo, sampleService sample.Service) {
	api.GET("/v1/samples/:referenceUUID", GetSampleByReferenceUUID(sampleService))
	api.POST("/v1/samples", CreateSample(sampleService))
}

// Sample godoc
// @Router /samples [POST]
// @Summary Request an Sample creation
// @Description Sample Description.
// @Tags Sample
// @Accept json
// @Param CreateSample body request.CreateSample{} true "Requested body for Sample Creation"
// @Produce json
// @Success 201 {object} response.Sample{} "Sample request successfully created"
// @Header 201 {string} Location "The endpoint location for the created ressource. E.g /samples/da74025d-07cf-4347-8e97-3073e83adc82"
// @Failure 400  {object}  response.Error
// @Failure 500  {object}  response.Error
func CreateSample(sampleService sample.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("creating Sample")

		var createSampleRequest request.CreateSample
		var createdSampleResponse response.Sample

		if err := c.Bind(&createSampleRequest); err != nil {
			log.Info().Interface("createSampleRequest", createSampleRequest).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		createDTO := sample.CreateDTO{
			ReferenceUUID: createSampleRequest.ReferenceUUID,
		}

		createdSample, err := sampleService.CreateSample(c.Request().Context(), createDTO)
		if err != nil {
			log.Error().Interface("createSampleRequest", createSampleRequest).Msg("error on Sample creation")
			status := http.StatusInternalServerError
			prob := problem.New(
				problem.Title("UNEXPECTED_ERROR"),
				problem.Detail("Internal Server Error"),
				problem.Status(status),
			)
			return c.JSON(status, prob)
		}

		log.Info().Msg("retrieving created sample")

		createdSampleResponse.FromSample(createdSample)
		c.Response().Header().Set("Location", fmt.Sprint(constants.SampleRessource, "/", createdSample.ReferenceUUID))
		return c.JSON(http.StatusCreated, createdSampleResponse)
	}
}

// GetSampleByReferenceUUID godoc
// @Router /v1/samples/:referenceUUID [GET]
// @Summary Request an Sample
// @Description In this endpoint you can request an sample
// @Tags Sample
// @Accept json
// @Param referenceUUID path string true "Reference"
// @Produce json
// @Success 200 {object} response.Sample{} "Sample requested"
// @Failure 404  {object}  response.Error "Sample not found"
// @Failure 500  {object}  response.Error
func GetSampleByReferenceUUID(sampleService sample.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		sampleResponse := response.Sample{}

		byReferenceUUID := c.Param("referenceUUID")

		fetchSample, err := sampleService.GetByReferenceUUID(c.Request().Context(), byReferenceUUID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Warn().Msgf("GetSampleByReferenceUUID not found by %s", byReferenceUUID)

				errResponse := problem.New(
					problem.Title(sample.ErrSampleNotFound.Error()),
					problem.Detail("Sample not found or invalid"),
					problem.Status(http.StatusNotFound),
				)

				return c.JSON(http.StatusNotFound, errResponse)
			}

			log.Error().Msgf("GetSampleByReferenceUUID error when trying by %s", byReferenceUUID)

			return c.JSON(http.StatusInternalServerError, constants.ProblemInternalServerError)
		}

		sampleResponse.FromSample(fetchSample)

		return c.JSON(http.StatusOK, sampleResponse)
	}
}

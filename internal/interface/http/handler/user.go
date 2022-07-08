package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ftec-project/internal/domain/user"
	"github.com/ftec-project/internal/infra/constants"
	"github.com/ftec-project/internal/interface/http/dto/request"
	"github.com/ftec-project/internal/interface/http/dto/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"schneider.vip/problem"
)

func MakeUserHandler(api *echo.Echo, userService user.Service) {
	api.GET("/v1/users/:referenceUUID", GetUserByReferenceUUID(userService))
	api.POST("/v1/users", CreateUser(userService))
}

// User godoc
// @Router /users [POST]
// @Summary Request an User creation
// @Description User Description.
// @Tags User
// @Accept json
// @Param CreateUser body request.CreateUser{} true "Requested body for User Creation"
// @Produce json
// @Success 201 {object} response.User{} "User request successfully created"
// @Header 201 {string} Location "The endpoint location for the created ressource. E.g /users/da74025d-07cf-4347-8e97-3073e83adc82"
// @Failure 400  {object}  response.Error
// @Failure 500  {object}  response.Error
func CreateUser(userService user.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("creating User")

		var createUserRequest request.CreateUser
		var createdUserResponse response.User

		if err := c.Bind(&createUserRequest); err != nil {
			log.Info().Interface("createUserRequest", createUserRequest).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		createDTO := user.CreateDTO{
			ReferenceUUID: createUserRequest.ReferenceUUID,
		}

		createdUser, err := userService.CreateUser(c.Request().Context(), createDTO)
		if err != nil {
			log.Error().Interface("createUserRequest", createUserRequest).Msg("error on User creation")
			status := http.StatusInternalServerError
			prob := problem.New(
				problem.Title("UNEXPECTED_ERROR"),
				problem.Detail("Internal Server Error"),
				problem.Status(status),
			)
			return c.JSON(status, prob)
		}

		log.Info().Msg("retrieving created user")

		createdUserResponse.FromUser(createdUser)
		c.Response().Header().Set("Location", fmt.Sprint(constants.UserRessource, "/", createdUser.ReferenceUUID))
		return c.JSON(http.StatusCreated, createdUserResponse)
	}
}

// GetUserByReferenceUUID godoc
// @Router /v1/users/:referenceUUID [GET]
// @Summary Request an User
// @Description In this endpoint you can request an user
// @Tags User
// @Accept json
// @Param referenceUUID path string true "Reference"
// @Produce json
// @Success 200 {object} response.User{} "User requested"
// @Failure 404  {object}  response.Error "User not found"
// @Failure 500  {object}  response.Error
func GetUserByReferenceUUID(userService user.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		userResponse := response.User{}

		byReferenceUUID := c.Param("referenceUUID")

		fetchUser, err := userService.GetByReferenceUUID(c.Request().Context(), byReferenceUUID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Warn().Msgf("GetUserByReferenceUUID not found by %s", byReferenceUUID)

				errResponse := problem.New(
					problem.Title(user.ErrUserNotFound.Error()),
					problem.Detail("User not found or invalid"),
					problem.Status(http.StatusNotFound),
				)

				return c.JSON(http.StatusNotFound, errResponse)
			}

			log.Error().Msgf("GetUserByReferenceUUID error when trying by %s", byReferenceUUID)

			return c.JSON(http.StatusInternalServerError, constants.ProblemInternalServerError)
		}

		userResponse.FromUser(fetchUser)

		return c.JSON(http.StatusOK, userResponse)
	}
}

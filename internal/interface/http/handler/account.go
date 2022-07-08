package handler

import (
	"fmt"
	"net/http"

	"github.com/ftec-project/internal/domain/account"
	"github.com/ftec-project/internal/infra/constants"
	"github.com/ftec-project/internal/interface/http/dto/request"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"schneider.vip/problem"
)

func MakeAccountHandler(api *echo.Echo, accountService account.Service) {
	api.POST("/v1/accounts", CreateAccount(accountService))
}

// Account godoc
// @Router /accounts [POST]
// @Summary Request an Account creation
// @Description Account Description.
// @Tags Account
// @Accept json
// @Param CreateAccount body request.CreateAccount{} true "Requested body for Account Creation"
// @Produce json
// @Success 201 {object} response.Account{} "Account request successfully created"
// @Header 201 {string} Location "The endpoint location for the created ressource. E.g /accounts/da74025d-07cf-4347-8e97-3073e83adc82"
// @Failure 400  {object}  response.Error
// @Failure 500  {object}  response.Error
func CreateAccount(accountService account.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("creating Account")

		var createAccountRequest request.CreateAccount
		var createdAccountResponse response.Account

		if err := c.Bind(&createAccountRequest); err != nil {
			log.Info().Interface("createAccountRequest", createAccountRequest).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		createDTO := account.CreateDTO{
			Email:    createAccountRequest.Email,
			Password: createAccountRequest.Password,
		}

		createdAccount, err := accountService.CreateAccount(c.Request().Context(), createDTO)
		if err != nil {
			log.Error().Interface("createAccountRequest", createAccountRequest).Msg("error on Account creation")
			status := http.StatusInternalServerError
			prob := problem.New(
				problem.Title("UNEXPECTED_ERROR"),
				problem.Detail("Internal Server Error"),
				problem.Status(status),
			)
			return c.JSON(status, prob)
		}

		log.Info().Msg("retrieving created account")

		createdAccountResponse.FromAccount(createdAccount)
		c.Response().Header().Set("Location", fmt.Sprint(constants.AccountRessource, "/", createdAccount.ReferenceUUID))
		return c.JSON(http.StatusCreated, createdAccountResponse)
	}
}

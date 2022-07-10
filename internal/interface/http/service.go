package api

import (
	_ "github.com/ftec-project/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ftec-project/internal/domain/account"
	"github.com/ftec-project/internal/domain/partner"
	"github.com/ftec-project/internal/domain/sample"
	"github.com/ftec-project/internal/domain/location"
	"github.com/ftec-project/internal/domain/user"
	"github.com/ftec-project/internal/interface/http/handler"
)

type serviceImpl struct {
	api  *echo.Echo
	Port string
}

func NewService(
	port string,
	applicationName string,
	sampleService sample.Service,
	accountService account.Service,
	userService user.Service,
	partnerService partner.Service,
	locationService location.Service,
) *serviceImpl {
	echoAPI := echo.New()

	handler.MakeHealtHandler(echoAPI)
	handler.MakeSampleHandler(echoAPI, sampleService)
	handler.MakeAccountHandler(echoAPI, accountService)
	handler.MakeUserHandler(echoAPI, userService)
	handler.MakePartnerHandler(echoAPI, partnerService)
	handler.MakeLocationHandler(echoAPI, locationService)

	echoAPI.GET("/swagger/*", echoSwagger.WrapHandler)
	return &serviceImpl{
		api:  echoAPI,
		Port: port,
	}
}

func (s serviceImpl) StartServer() error {
	return s.api.Start(s.Port)
}

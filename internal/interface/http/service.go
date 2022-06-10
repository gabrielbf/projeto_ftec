package api

import (
	_ "github.com/sample-project/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/sample-project/internal/domain/sample"
	"github.com/sample-project/internal/interface/http/handler"
)

type serviceImpl struct {
	api  *echo.Echo
	Port string
}

func NewService(
	port string,
	applicationName string,
	sampleService sample.Service,
) *serviceImpl {
	echoAPI := echo.New()

	handler.MakeHealtHandler(echoAPI)
	handler.MakeSampleHandler(echoAPI, sampleService)

	echoAPI.GET("/swagger/*", echoSwagger.WrapHandler)
	return &serviceImpl{
		api:  echoAPI,
		Port: port,
	}
}

func (s serviceImpl) StartServer() error {
	return s.api.Start(s.Port)
}

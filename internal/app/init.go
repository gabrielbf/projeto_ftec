package app

import (
	"fmt"
	"path/filepath"

	"github.com/ftec-project/internal/app/service"
	"github.com/ftec-project/internal/infra/constants"
	"github.com/ftec-project/internal/infra/database"
	"github.com/ftec-project/internal/infra/database/repository"
	"github.com/ftec-project/internal/infra/environment"
	logConfig "github.com/ftec-project/internal/infra/log"
	api "github.com/ftec-project/internal/interface/http"
	"github.com/rs/zerolog/log"
)

type TracerLogger struct {
}

func (t TracerLogger) Log(msg string) {
	log.Info().Msg(msg)
}

func Init(rootdir string) {
	fmt.Println(rootdir)
	environment.LoadEnv(filepath.Join(rootdir, ".env"))
	logConfig.ConfigZeroLog(environment.Env.LogLevel)

	pgService, err := database.NewPostgres(database.Config{
		Host:     environment.Env.DbHost,
		Port:     environment.Env.DbPort,
		User:     environment.Env.DbUser,
		Password: environment.Env.DbPassword,
		Database: environment.Env.DbName,
	})

	if err != nil {
		log.Fatal().Stack().Err(err).Send()
	}

	migrationsPath := "file://" + filepath.Join(rootdir, "internal/infra/database/migrations")
	err = database.MigrateUp(pgService, environment.Env.DbName, migrationsPath)
	if err != nil {
		log.Fatal().Stack().Err(err).Send()
	}

	sampleRepository := repository.NewSampleRepository(pgService)
	contactRepository := repository.NewContactRepository(pgService)
	accountRepository := repository.NewAccountRepository(pgService)
	userRepository := repository.NewUserRepository(pgService)
	partnerRepository := repository.NewPartnerRepository(pgService)

	sampleService := service.NewSampleService(
		sampleRepository,
	)

	contactService := service.NewContactService(
		contactRepository,
	)

	userService := service.NewUserService(
		userRepository,
	)

	partnerService := service.NewPartnerService(
		partnerRepository,
	)

	accountService := service.NewAccountService(
		accountRepository,
		contactService,
		userService,
		partnerService,
	)

	apiService := api.NewService(fmt.Sprintf(":%s", environment.Env.HttpPort), constants.ServiceName,
		sampleService,
		accountService,
		userService,
		partnerService,
	)

	go func() {
		if err := apiService.StartServer(); err != nil {
			log.Fatal().Stack().Err(err).Send()
		}
	}()
}

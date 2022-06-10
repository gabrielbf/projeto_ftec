package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lib/pq"
	"github.com/sample-project/internal/infra/constants"

	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	gormpostgres "gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
	)
}

func NewPostgres(config Config) (*gorm.DB, error) {
	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
	}

	sqltrace.Register("postgres", &pq.Driver{})

	tracedb, err := sqltrace.Open("postgres", config.String(), sqltrace.WithServiceName(constants.ServiceName))

	if err != nil {
		return nil, err
	}

	db, err := gormtrace.Open(
		gormpostgres.New(
			gormpostgres.Config{
				Conn: tracedb,
			},
		),
		&gormConfig,
		gormtrace.WithServiceName(constants.ServiceName),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateUp(gormInstance *gorm.DB, database string, path string) error {
	sqlInstance, err := gormInstance.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(sqlInstance, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		path,
		database,
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

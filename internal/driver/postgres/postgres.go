package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/arvinpaundra/ngekost-api/pkg/util/config"
	l "github.com/arvinpaundra/ngekost-api/pkg/util/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
	sslmode  string
	timezone string
}

func New() *Postgres {
	return &Postgres{
		user:     config.GetString("POSTGRES_USER"),
		password: config.GetString("POSTGRES_PASSWORD"),
		host:     config.GetString("POSTGRES_HOST"),
		port:     config.GetString("POSTGRES_PORT"),
		dbname:   config.GetString("POSTGRES_DBNAME"),
		sslmode:  config.GetString("POSTGRES_SSLMODE"),
		timezone: config.GetString("POSTGRES_TIMEZONE"),
	}
}

func (p *Postgres) Connect(ctx context.Context) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		p.host,
		p.user,
		p.password,
		p.dbname,
		p.port,
		p.sslmode,
		p.timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Logging().Error(err.Error())
		os.Exit(1)
	}

	log.Println("connected to postgres")

	return db
}

package infraestructure

import (
	"errors"
	"fmt"
	"github.com/DaemonGearIT/go-sample-rest/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgres struct {
	Database *gorm.DB
}

var database *gorm.DB

func (p *Postgres) Connect(conf config.Configuration) error {
	if &conf == nil {
		return errors.New("Configuration must not be nil")
	}

	if &(conf.DATABASE) == nil {
		return errors.New("Configuration DATABASE must not be nil, please verify config.yml")
	}

	dialect := conf.DATABASE.DIALECT
	url := fmt.Sprintf("host=%v user=%v dbname=%v sslmode=disable password=%v", conf.DATABASE.HOST, conf.DATABASE.USER, conf.DATABASE.DBNAME, conf.DATABASE.PASSWORD)
	db, err := gorm.Open(dialect, url)
	if err != nil {
		return err
	}

	database = db
	database.LogMode(true)

	return nil
}

func (p *Postgres) Migrate(values ...interface{}) error {
	if &(p.Database) == nil {
		return errors.New("Database must not be nul, configure first")
	}

	p.Database.AutoMigrate(values)
	return nil
}

func PostgresSession() *Postgres {
	return &Postgres{database}
}

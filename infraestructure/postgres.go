package infraestructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgres struct {
	Database gorm.DB
}

var database gorm.DB

func (p *Postgres) Connect() error {
	database, err = gorm.Open("postgres", "host=localhost user=daemongear dbname=daemongear sslmode=disable password=daemongear")

	if err != nil {
		fmt.Printf("Failed connect to database: %v", err.Error())
	}
}

func PostgresSession() *Postgres {
	return &Postgres{database.Copy()}
}

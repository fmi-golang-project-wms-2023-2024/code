package postgres

import (
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type serviceDB struct {
	*gorm.DB
}

func MustNew(dsn string) *serviceDB {
	// Quick hacky fix to wait for the db to become alive before interacting with it.
	time.Sleep(time.Second * 5)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &serviceDB{DB: db}
}

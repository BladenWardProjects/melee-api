package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/BladenWard/melee-api/types"
)

type DB struct {
	*gorm.DB
}

func Init() *DB {
	db, err := gorm.Open(postgres.Open("host=localhost port=65432 user=postgres password=postgres dbname=postgres sslmode=disable TimeZone=UTC"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")

	err = db.AutoMigrate(
		&types.Character{},
		&types.GroundAttack{},
		&types.Aerial{},
		&types.Special{},
		&types.Grab{},
		&types.Throw{},
		&types.Dodge{},

		&types.Song{},
	)
	if err != nil {
		panic("failed to auto migrate: " + err.Error())
	}

	return &DB{db}
}

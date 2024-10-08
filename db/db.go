package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/BladenWard/melee-api/types"
)

type DB struct {
	*gorm.DB
}

func Init() *DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")
	db, err := gorm.Open(postgres.Open("host=localhost port=65432 user="+db_user+" password="+db_pass+" dbname="+db_name+" sslmode=disable TimeZone=UTC"), &gorm.Config{
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

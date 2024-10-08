package seed

import (
	"fmt"

	"github.com/BladenWard/melee-api/db"
)

func Seed(db *db.DB) {
	fmt.Println("Seeding database...")

	DROP_TABLES(db)

	seedCharacters(db)
	seedSongs(db)

	fmt.Println("Finished seeding database.")
}

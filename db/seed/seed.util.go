package seed

import (
	"github.com/BladenWard/melee-api/db"
	"github.com/BladenWard/melee-api/types"
)

func DROP_TABLES(db *db.DB) {
	err := db.DB.Migrator().DropTable(
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
		panic("failed to drop tables: " + err.Error())
	}
	db.DB.AutoMigrate(
		&types.Character{},
		&types.GroundAttack{},
		&types.Aerial{},
		&types.Special{},
		&types.Grab{},
		&types.Throw{},
		&types.Dodge{},
		&types.Song{},
	)
}

package seed

import (
	"fmt"
	"os"

	"github.com/BladenWard/melee-api/db"
	"github.com/BladenWard/melee-api/types"
)

var charList = []string{
	"bowser",
	"captain_falcon",
	"donkey_kong",
	"dr._mario",
	"falco",
	"fox",
	"mr._game_&_watch",
	"ganondorf",
	"ice_climbers",
	"kirby",
	"jigglypuff",
	"link",
	"luigi",
	"mario",
	"marth",
	"mewtwo",
	"ness",
	"peach",
	"pichu",
	"pikachu",
	"roy",
	"samus",
	"sheik",
	"yoshi",
	"young_link",
	"zelda",
}

func seedCharacter(db *db.DB, character *types.Character, id int) {
	characterFile, _ := os.ReadFile("seed/characters/" + charList[id] + ".json")
	characterJson := string(characterFile)

	types.SeedCharacterStructure(id, character, &characterJson)

	fmt.Println(charList[id] + " seeded. ID: " + fmt.Sprint(character.ID))
	db.DB.Create(&character)
}

func SeedCharacters(db *db.DB) {
	char := types.Character{}
	for i := 0; i < 5; i++ {
		fmt.Println("Seeding character " + charList[i] + "...")
		seedCharacter(db, &char, i)
	}
	// seedCharacter(db, &char, 0)

	newChar := types.Character{}
	db.First(&newChar, 1)
	fmt.Println(newChar.GetMoveByName("up_air"))
	fmt.Println(newChar.Name)
}

func DROP_TABLES(db *db.DB) {
	err := db.DB.Migrator().DropTable(
		&types.Character{},
		&types.GroundAttack{},
		&types.Aerial{},
		&types.Special{},
		&types.Grab{},
		&types.Throw{},
		&types.Dodge{},
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
	)
}

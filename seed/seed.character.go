package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BladenWard/melee-api/db"
	"github.com/BladenWard/melee-api/types"
)

func seedCharacterStructure(character *types.Character, characterJson *string) {
	infoMap := map[string]interface{}{}
	json.Unmarshal([]byte(*characterJson), &infoMap)

	seedStats(0, character, infoMap)
	seedGroundAttacks(character, infoMap["ground"].([]interface{}))
	seedAerials(character, infoMap["aerial"].([]interface{}))
	seedSpecials(character, infoMap["special"].([]interface{}))
	seedGrabs(character, infoMap["grab"].([]interface{}))
	seedThrows(character, infoMap["throw"].([]interface{}))
	seedDodges(character, infoMap["dodge"].([]interface{}))
}

var charList = []string{
	"bowser",
	"captain_falcon",
	"donkey_kong",
	"dr._mario",
	"falco",
	"fox",
	"mr._game_and_watch",
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

func seedCharacter(db *db.DB, id uint) {
	characterFile, _ := os.ReadFile("seed/characters/" + charList[id] + ".json")
	characterJson := string(characterFile)

	character := types.Character{}
	seedCharacterStructure(&character, &characterJson)

	db.DB.Create(&character)
}

func SeedCharacters(db *db.DB) {
	seedCharacter(db, 1)
	// seedCharacter(db, 2)

	char := types.Character{}
	db.First(&char)
	fmt.Println(char.Name)
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

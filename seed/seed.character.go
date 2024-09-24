package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BladenWard/melee-api/db"
	"github.com/BladenWard/melee-api/types"
)

func seedCharacterStructure(charId int, character *types.Character, characterJson *string) {
	infoMap := map[string]interface{}{}
	json.Unmarshal([]byte(*characterJson), &infoMap)

	seedStats(charId, character, infoMap)
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
	fmt.Println("Seeding character " + charList[id] + "...")
	characterJson := string(characterFile)

	seedCharacterStructure(id, character, &characterJson)

	db.DB.Create(&character)
}

func SeedCharacters(db *db.DB) {
	char := types.Character{}
	for i := 0; i < len(charList); i++ {
		seedCharacter(db, &char, i)
	}
	// seedCharacter(db, &char, 0)

	newChar := types.Character{}
	db.First(&newChar)
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

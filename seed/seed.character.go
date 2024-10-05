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

func seedCharacters(db *db.DB) {
	for i := 0; i < len(charList); i++ {
		char := types.Character{}
		fmt.Println("Seeding character " + charList[i] + "...")
		seedCharacter(db, &char, i)
		fmt.Println(charList[i] + " seeded. ID: " + fmt.Sprint(char.ID))
	}
}

func seedCharacter(db *db.DB, character *types.Character, id int) {
	characterFile, err := os.ReadFile("seed/characters/" + charList[id] + ".json")
	if err != nil {
		panic(err)
	}
	characterJson := string(characterFile)

	character.SerializeStructure(id, &characterJson)

	db.Create(&character)
}

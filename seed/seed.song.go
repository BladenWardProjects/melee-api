package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BladenWard/melee-api/db"
	"github.com/BladenWard/melee-api/types"
)

func seedSongs(db *db.DB) error {
	fmt.Println("Seeding songs...")

	songJson, err := os.ReadFile("seed/songs/songs.json")
	if err != nil {
		return err
	}

	songMap := map[string]interface{}{}
	err = json.Unmarshal([]byte(songJson), &songMap)
	if err != nil {
		return err
	}

	songs := songMap["songs"].([]interface{})

	for id, song := range songs {
		db_song := types.Song{
			ID:          uint(id + 1),
			Title:       song.(map[string]interface{})["title"].(string),
			Origin:      song.(map[string]interface{})["origin"].(string),
			PlaysDuring: song.(map[string]interface{})["plays_during"].(string),
		}

		fmt.Println("Seeding song " + db_song.Title + "...")
		if err := db.Create(&db_song).Error; err != nil {
			return err
		}
	}

	return nil
}

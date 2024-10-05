package db

import (
	"github.com/BladenWard/melee-api/types"
)

func (db *DB) GetCharacterByID(id uint) (types.Character, error) {
	var character types.Character
	err := db.First(&character, id).Error
	return character, err
}

func (db *DB) GetCharacterByName(name string) (types.Character, error) {
	var character types.Character
	err := db.Where("name = ?", name).First(&character).Error
	return character, err
}

func (db *DB) GetCharacters() ([]types.Character, error) {
	var characters []types.Character
	err := db.Find(&characters).Error
	return characters, err
}

func (db *DB) GetSongs() ([]types.Song, error) {
	var songs []types.Song
	err := db.Find(&songs).Error
	return songs, err
}

func (db *DB) GetSongByID(id uint) (types.Song, error) {
	var song types.Song
	err := db.First(&song, id).Error
	return song, err
}

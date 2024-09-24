package db

import "github.com/BladenWard/melee-api/types"

type Storage interface {
	GetCharacterByID(id uint) (types.Character, error)
	GetCharacterByName(name string) (types.Character, error)
	GetCharacters() ([]types.Character, error)
}

func (db *DB) GetCharacterByID(id uint) (types.Character, error) {
	var character types.Character
	err := db.First(&character, id).Error
	return character, err
}

func (db *DB) GetCharacterByName(name string) (types.Character, error) {
	var character types.Character
	err := db.Where("name = ?", name).Select("id, name").Find(&character).Error
	return character, err
}

func (db *DB) GetCharacters() ([]types.Character, error) {
	var characters []types.Character
	err := db.Find(&characters).Error
	return characters, err
}

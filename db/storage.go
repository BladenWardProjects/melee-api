package db

import "github.com/BladenWard/melee-api/types"

type Storage interface {
	GetCharacterByID(id uint) (types.Character, error)
	GetCharacterByName(name string) (types.Character, error)
	GetCharacters() ([]types.Character, error)
	GetSongs() ([]types.Song, error)
	GetSongByID(id uint) (types.Song, error)
}

package types

import "database/sql"

type Character struct {
	ID             uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	Name           string         `json:"name"`
	Weight         uint           `json:"weight"`
	FastfallSpeed  float64        `json:"fastfall_speed"`
	DashSpeed      float64        `json:"dash_speed"`
	RunSpeed       float64        `json:"run_speed"`
	WavedashLength uint           `json:"wavedash_length_rank"`
	Galint         uint           `json:"galint"`
	JumpSquat      uint           `json:"jump_squat"`
	Walljump       bool           `json:"walljump"`
	GroundAttacks  []GroundAttack `json:"ground" gorm:"serializer:json;foreignKey:CharacterID"`
	Aerials        []Aerial       `json:"aerial" gorm:"serializer:json;foreignKey:CharacterID"`
	Specials       []Special      `json:"special" gorm:"serializer:json;foreignKey:CharacterID"`
	Grabs          []Grab         `json:"grab" gorm:"serializer:json;foreignKey:CharacterID"`
	Throws         []Throw        `json:"throw" gorm:"serializer:json;foreignKey:CharacterID"`
	Dodges         []Dodge        `json:"dodge" gorm:"serializer:json;foreignKey:CharacterID"`
}

type GroundAttack struct {
	ID          uint          `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Name        string        `json:"name"`
	Start       sql.NullInt16 `json:"start"`
	End         uint          `json:"end"`
	TotalFrames *uint         `json:"frames"`
	IasaFrames  *uint         `json:"iasa_frames"`
	ShieldStun  sql.NullInt16 `json:"shield_stun"`
	BaseDamage  uint          `json:"base_damage"`
	WeakDamage  *uint         `json:"weak_damage"`
	CharacterID uint
}

type Aerial struct {
	ID          uint          `json:"id,omitempty" gorm:"primaryKey,unique"`
	Name        string        `json:"name"`
	Start       uint          `json:"start"`
	End         *uint         `json:"end"`
	TotalFrames *uint         `json:"frames"`
	ShieldStun  uint          `json:"shield_stun"`
	BaseDamage  uint          `json:"base_damage"`
	WeakDamage  *uint         `json:"weak_damage"`
	LandingLag  uint          `json:"landing_lag"`
	LCancelLag  sql.NullInt16 `json:"l_cancel_lag"`
	AutoCancel  *int          `json:"auto_cancel"`
	CharacterID uint
}

type Special struct {
	ID                 uint   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Name               string `json:"name"`
	Start              uint   `json:"start"`
	End                *uint  `json:"end"`
	TotalFrames        *uint  `json:"frames"`
	ShieldStun         *uint  `json:"shield_stun"`
	BaseDamage         uint   `json:"base_damage"`
	WeakDamage         *uint  `json:"weak_damage"`
	LandingLag         *uint  `json:"landing_lag"`
	LandingFallSpecial *uint  `json:"landing_fall_special,omitempty"`
	CharacterID        uint
}

type Grab struct {
	ID          uint   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Name        string `json:"name"`
	Start       *uint  `json:"start"`
	TotalFrames uint   `json:"frames"`
	CharacterID uint
}

type Throw struct {
	ID          uint   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Name        string `json:"name"`
	Start       *uint  `json:"start"`
	End         *uint  `json:"end"`
	TotalFrames uint   `json:"frames"`
	BaseDamage  uint   `json:"base_damage"`
	WeakDamage  *uint  `json:"weak_damage"`
	CharacterID uint
}

type Dodge struct {
	ID                 uint   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Name               string `json:"name"`
	Start              uint   `json:"start"`
	End                uint   `json:"end"`
	TotalFrames        uint   `json:"frames"`
	LandingFallSpecial *uint  `json:"landing_fall_special,omitempty"`
	CharacterID        uint
}

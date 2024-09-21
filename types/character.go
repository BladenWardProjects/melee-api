package types

import (
	"context"
	"errors"
	"reflect"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm/schema"
)

type YamlMap map[string]map[string]interface{}

func (ym *YamlMap) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error {
	if dbValue == nil {
		*ym = nil
		return nil
	}
	bytes, ok := dbValue.([]byte)
	if !ok {
		return errors.New("failed to unmarshal YAML value: source data is not []byte")
	}

	var m map[string]map[string]interface{}
	if err := yaml.Unmarshal(bytes, &m); err != nil {
		return err
	}

	*ym = m
	return nil
}

// NOTE: Refactor this, or remove it if i dont use it
func (ym *YamlMap) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	if ym == nil {
		return nil, nil
	}
	return yaml.Marshal(ym)
}

// TODO: Fix script, move the damage array into 2 fields
type Character struct {
	Name           string         `json:"name"`
	ID             uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	Weight         uint           `json:"weight"`
	FastfallSpeed  float64        `json:"fastfall_speed"`
	DashSpeed      float64        `json:"dash_speed"`
	RunSpeed       float64        `json:"run_speed"`
	WavedashLength uint           `json:"wavedash_length_rank"`
	Galint         uint           `json:"galint"`
	JumpSquat      uint           `json:"jump_squat"`
	Walljump       bool           `json:"walljump"`
	GroundAttacks  []GroundAttack `json:"ground" gorm:"serializer:json"`
	Aerials        []Aerial       `json:"aerial" gorm:"serializer:json"`
	Specials       []Special      `json:"special" gorm:"serializer:json"`
	Grabs          []Grab         `json:"grab" gorm:"serializer:json"`
	Throws         []Throw        `json:"throw" gorm:"serializer:json"`
	Dodges         []Dodge        `json:"dodge" gorm:"serializer:json"`
}

type GroundAttack struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Start       uint   `json:"start"`
	End         uint   `json:"end"`
	TotalFrames uint   `json:"frames"`
	IasaFrames  *uint  `json:"iasa_frames"`
	ShieldStun  uint   `json:"shield_stun"`
	BaseDamage  uint   `json:"base_damage"`
	WeakDamage  *uint  `json:"weak_damage"`
	CharacterID uint
}

type Aerial struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Start       uint   `json:"start"`
	End         uint   `json:"end"`
	TotalFrames uint   `json:"frames"`
	ShieldStun  uint   `json:"shield_stun"`
	BaseDamage  uint   `json:"base_damage"`
	WeakDamage  *uint  `json:"weak_damage"`
	LandingLag  uint   `json:"landing_lag"`
	LCancelLag  uint   `json:"l_cancel_lag"`
	AutoCancel  *int   `json:"auto_cancel"`
	CharacterID uint
}

type Special struct {
	ID                 uint   `json:"id" gorm:"primaryKey"`
	Name               string `json:"name"`
	Start              uint   `json:"start"`
	End                uint   `json:"end"`
	TotalFrames        uint   `json:"frames"`
	ShieldStun         *uint  `json:"shield_stun"`
	BaseDamage         uint   `json:"base_damage"`
	WeakDamage         *uint  `json:"weak_damage"`
	LandingLag         *uint  `json:"landing_lag"`
	LandingFallSpecial *uint  `json:"landing_fall_special"`
	CharacterID        uint
}

type Grab struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Start       uint   `json:"start"`
	TotalFrames uint   `json:"frames"`
	CharacterID uint
}

type Throw struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Start       uint   `json:"start"`
	End         uint   `json:"end"`
	TotalFrames uint   `json:"frames"`
	BaseDamage  uint   `json:"base_damage"`
	WeakDamage  *uint  `json:"weak_damage"`
	CharacterID uint
}

type Dodge struct {
	ID                 uint   `json:"id" gorm:"primaryKey"`
	Name               string `json:"name"`
	Start              uint   `json:"start"`
	End                uint   `json:"end"`
	TotalFrames        uint   `json:"frames"`
	LandingFallSpecial *uint  `json:"landing_fall_special"`
	CharacterID        uint
}

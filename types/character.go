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
	Name           string         `yaml:"name"`
	ID             uint           `yaml:"id" gorm:"primaryKey,autoIncrement"`
	Weight         uint           `yaml:"weight"`
	FastfallSpeed  float64        `yaml:"fastfall_speed"`
	DashSpeed      float64        `yaml:"dash_speed"`
	RunSpeed       float64        `yaml:"run_speed"`
	WavedashLength uint           `yaml:"wavedash_length_rank"`
	Galint         uint           `yaml:"galint"`
	JumpSquat      uint           `yaml:"jump_squat"`
	Walljump       bool           `yaml:"walljump"`
	GroundAttacks  []GroundAttack `yaml:"ground"`
	Aerials        []Aerial       `yaml:"aerial"`
	Specials       []Special      `yaml:"special"`
	Grabs          []Grab         `yaml:"grab"`
	Throws         []Throw        `yaml:"throw"`
	Dodges         []Dodge        `yaml:"dodge"`
}

type GroundAttack struct {
	ID          uint   `yaml:"id" gorm:"primaryKey"`
	Name        string `yaml:"name"`
	Start       uint   `yaml:"start"`
	End         uint   `yaml:"end"`
	TotalFrames uint   `yaml:"total_frames"`
	IasaFrames  *uint  `yaml:"iasa_frames"`
	ShieldStun  uint   `yaml:"shield_stun"`
	BaseDamage  uint32 `yaml:"base_damage"`
	WeakDamage  uint32 `yaml:"weak_damage"`
	CharacterID uint
}

type Aerial struct {
	ID          uint   `yaml:"id" gorm:"primaryKey"`
	Name        string `yaml:"name"`
	Start       uint   `yaml:"start"`
	End         uint   `yaml:"end"`
	TotalFrames uint   `yaml:"total_frames"`
	ShieldStun  uint   `yaml:"shield_stun"`
	BaseDamage  uint32 `yaml:"base_damage"`
	WeakDamage  uint32 `yaml:"weak_damage"`
	LandingLag  uint   `yaml:"landing_lag"`
	LCancelLag  uint   `yaml:"l_cancel_lag"`
	AutoCancel  *int   `yaml:"auto_cancel"`
	CharacterID uint
}

type Special struct {
	ID                 uint   `yaml:"id" gorm:"primaryKey"`
	Name               string `yaml:"name"`
	Start              uint   `yaml:"start"`
	End                uint   `yaml:"end"`
	TotalFrames        uint   `yaml:"total_frames"`
	ShieldStun         *uint  `yaml:"shield_stun"`
	BaseDamage         uint32 `yaml:"base_damage"`
	WeakDamage         uint32 `yaml:"weak_damage"`
	LandingLag         *uint  `yaml:"landing_lag"`
	LandingFallSpecial *uint  `yaml:"landing_fall_special"`
	CharacterID        uint
}

type Grab struct {
	ID          uint   `yaml:"id" gorm:"primaryKey"`
	Name        string `yaml:"name"`
	Start       uint   `yaml:"start"`
	TotalFrames uint   `yaml:"total_frames"`
	CharacterID uint
}

type Throw struct {
	ID          uint   `yaml:"id" gorm:"primaryKey"`
	Name        string `yaml:"name"`
	Start       uint   `yaml:"start"`
	End         uint   `yaml:"end"`
	TotalFrames uint   `yaml:"total_frames"`
	BaseDamage  uint32 `yaml:"base_damage"`
	WeakDamage  uint32 `yaml:"weak_damage"`
	CharacterID uint
}

type Dodge struct {
	ID                 uint   `yaml:"id" gorm:"primaryKey"`
	Name               string `yaml:"name"`
	Start              uint   `yaml:"start"`
	End                uint   `yaml:"end"`
	TotalFrames        uint   `yaml:"total_frames"`
	LandingFallSpecial *uint  `yaml:"landing_fall_special"`
	CharacterID        uint
}

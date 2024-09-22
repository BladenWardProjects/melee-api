package types

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

type Move interface {
	GetMoveByName(string) interface{}
}

func (c *Character) GetMoveByName(name string) interface{} {
	for _, move := range c.GroundAttacks {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Aerials {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Specials {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Grabs {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Throws {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Dodges {
		if move.Name == name {
			return move
		}
	}
	return nil
}

type GroundAttack struct {
	ID          uint   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Name        string `json:"name"`
	Start       uint   `json:"start"`
	End         uint   `json:"end"`
	TotalFrames *uint  `json:"frames"`
	IasaFrames  *uint  `json:"iasa_frames"`
	ShieldStun  uint   `json:"shield_stun"`
	BaseDamage  uint   `json:"base_damage"`
	WeakDamage  *uint  `json:"weak_damage"`
	CharacterID uint
}

type Aerial struct {
	ID          uint   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Name        string `json:"name"`
	Start       uint   `json:"start"`
	End         *uint  `json:"end"`
	TotalFrames *uint  `json:"frames"`
	ShieldStun  uint   `json:"shield_stun"`
	BaseDamage  uint   `json:"base_damage"`
	WeakDamage  *uint  `json:"weak_damage"`
	LandingLag  uint   `json:"landing_lag"`
	LCancelLag  uint   `json:"l_cancel_lag"`
	AutoCancel  *int   `json:"auto_cancel"`
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
	LandingFallSpecial *uint  `json:"landing_fall_special"`
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
	LandingFallSpecial *uint  `json:"landing_fall_special"`
	CharacterID        uint
}

package types

type Character struct {
	Name           string          `yaml:"name"`
	ID             uint            `yaml:"id" gorm:"primaryKey"`
	Weight         uint            `yaml:"weight"`
	FastfallSpeed  float64         `yaml:"fastfall_speed"`
	DashSpeed      float64         `yaml:"dash_speed"`
	RunSpeed       float64         `yaml:"run_speed"`
	WavedashLength uint            `yaml:"wavedash_length"`
	Galint         uint            `yaml:"galinth_length"`
	JumpSquat      uint            `yaml:"jump_squat"`
	Walljump       bool            `yaml:"walljump"`
	Moves          map[string]Move `yaml:"moves"`
	// TODO: Add moves to the character that might be a big issue
	// GroundAttacks []GroundAttack `yaml:"ground"`
	// Aerials       []Aerial       `yaml:"aerial"`
	// Specials      []Special      `yaml:"special"`
	// Grabs         []Grab         `yaml:"grab"`
	// Throws []Throw `yaml:"throw"`
	// Dodges []Dodge `yaml:"dodge"`
}

// TODO: Create a type for generic moves
type Move struct {
	ID uint `yaml:"id" gorm:"primaryKey"`

	// GroundAttacks []GroundAttack `yaml:"ground"`
	// Aerials       []Aerial       `yaml:"aerial"`
	// Specials      []Special      `yaml:"special"`
	// Grabs         []Grab         `yaml:"grab"`
	// Throws        []Throw        `yaml:"throw"`
	// Dodges        []Dodge        `yaml:"dodge"`
	CharacterID uint
}

type GroundAttack struct {
	ID          uint     `yaml:"id" gorm:"primaryKey"`
	Name        string   `yaml:"name"`
	Start       uint     `yaml:"start"`
	End         uint     `yaml:"end"`
	TotalFrames uint     `yaml:"total_frames"`
	IasaFrames  *uint    `yaml:"iasa_frames"`
	ShieldStun  uint     `yaml:"shield_stun"`
	Damages     []Damage `yaml:"damages" gorm:"foreignKey:DamageNumbersID"`
	MoveID      uint
}

type Aerial struct {
	ID          uint     `yaml:"id" gorm:"primaryKey"`
	Name        string   `yaml:"name"`
	Start       uint     `yaml:"start"`
	End         uint     `yaml:"end"`
	TotalFrames uint     `yaml:"total_frames"`
	ShieldStun  uint     `yaml:"shield_stun"`
	Damages     []Damage `yaml:"damages" gorm:"foreignKey:DamageNumbersID"`
	LandingLag  uint     `yaml:"landing_lag"`
	LCancelLag  uint     `yaml:"l_cancel_lag"`
	AutoCancel  *int     `yaml:"auto_cancel"`
	MoveID      uint
}

type Special struct {
	ID                 uint     `yaml:"id" gorm:"primaryKey"`
	Name               string   `yaml:"name"`
	Start              uint     `yaml:"start"`
	End                uint     `yaml:"end"`
	TotalFrames        uint     `yaml:"total_frames"`
	ShieldStun         *uint    `yaml:"shield_stun"`
	Damages            []Damage `yaml:"damages"`
	LandingLag         *uint    `yaml:"landing_lag"`
	LandingFallSpecial *uint    `yaml:"landing_fall_special"`
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
	ID          uint     `yaml:"id" gorm:"primaryKey"`
	Name        string   `yaml:"name"`
	Start       uint     `yaml:"start"`
	End         uint     `yaml:"end"`
	TotalFrames uint     `yaml:"total_frames"`
	Damages     []Damage `yaml:"damages" gorm:"foreignKey:MoveID"`
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

type Damage struct {
	ID uint `yaml:"id" gorm:"primaryKey"`
	// MoveID int  `yaml:"move_id"`
	DamageNumbersID uint `yaml:"damage_numbers_id"`
	Number          int  `yaml:"number"`
}

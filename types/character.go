package types

type Character struct {
	Name           string  `yaml:"name"`
	ID             uint    `yaml:"id" gorm:"primaryKey,autoIncrement"`
	Weight         uint    `yaml:"weight"`
	FastfallSpeed  float64 `yaml:"fastfall_speed"`
	DashSpeed      float64 `yaml:"dash_speed"`
	RunSpeed       float64 `yaml:"run_speed"`
	WavedashLength uint    `yaml:"wavedash_length_rank"`
	Galint         uint    `yaml:"galint"`
	JumpSquat      uint    `yaml:"jump_squat"`
	Walljump       bool    `yaml:"walljump"`
	// TODO: Serialize the moves
	// GroundAttacks []GroundAttack `yaml:"ground"`
	// Aerials        []Aerial       `yaml:"aerials"`
	// Specials       []Special      `yaml:"specials"`
	// Grabs          []Grab         `yaml:"grabs"`
	// Throws         []Throw        `yaml:"throws"`
	// Dodges         []Dodge        `yaml:"dodges"`
}

type GroundAttack struct {
	ID          uint    `yaml:"id" gorm:"primaryKey"`
	Name        string  `yaml:"name"`
	Start       uint    `yaml:"start"`
	End         uint    `yaml:"end"`
	TotalFrames uint    `yaml:"total_frames"`
	IasaFrames  *uint   `yaml:"iasa_frames"`
	ShieldStun  uint    `yaml:"shield_stun"`
	Damages     *[]uint `yaml:"damages"`
	CharacterID uint
}

type Aerial struct {
	ID          uint    `yaml:"id" gorm:"primaryKey"`
	Name        string  `yaml:"name"`
	Start       uint    `yaml:"start"`
	End         uint    `yaml:"end"`
	TotalFrames uint    `yaml:"total_frames"`
	ShieldStun  uint    `yaml:"shield_stun"`
	Damages     *[]uint `yaml:"damages"`
	LandingLag  uint    `yaml:"landing_lag"`
	LCancelLag  uint    `yaml:"l_cancel_lag"`
	AutoCancel  *int    `yaml:"auto_cancel"`
	CharacterID uint
}

type Special struct {
	ID                 uint    `yaml:"id" gorm:"primaryKey"`
	Name               string  `yaml:"name"`
	Start              uint    `yaml:"start"`
	End                uint    `yaml:"end"`
	TotalFrames        uint    `yaml:"total_frames"`
	ShieldStun         *uint   `yaml:"shield_stun"`
	Damages            *[]uint `yaml:"damages"`
	LandingLag         *uint   `yaml:"landing_lag"`
	LandingFallSpecial *uint   `yaml:"landing_fall_special"`
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
	ID          uint    `yaml:"id" gorm:"primaryKey"`
	Name        string  `yaml:"name"`
	Start       uint    `yaml:"start"`
	End         uint    `yaml:"end"`
	TotalFrames uint    `yaml:"total_frames"`
	Damages     *[]uint `yaml:"damages"`
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

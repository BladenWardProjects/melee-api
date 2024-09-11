package types

// TODO: add gorm tags and associations
type Character struct {
	Name           string  `json:"name"`
	ID             uint    `json:"id" gorm:"primaryKey"`
	Weight         uint    `json:"weight"`
	FastfallSpeed  float64 `json:"fastfall_speed"`
	DashSpeed      float64 `json:"dash_speed"`
	RunSpeed       float64 `json:"run_speed"`
	WavedashLength uint    `json:"wavedash_length"`
	Galint         uint    `json:"galinth_length"`
	JumpSquat      uint    `json:"jump_squat"`
	Walljump       bool    `json:"walljump"`
}

type GroundAttack struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Start       uint     `json:"start"`
	End         uint     `json:"end"`
	TotalFrames uint     `json:"total_frames"`
	IasaFrames  *uint    `json:"iasa_frames"`
	ShieldStun  uint     `json:"shield_stun"`
	Damages     []Damage `json:"damages"`
}

type Aerial struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Start       uint     `json:"start"`
	End         uint     `json:"end"`
	TotalFrames uint     `json:"total_frames"`
	ShieldStun  uint     `json:"shield_stun"`
	Damages     []Damage `json:"damages"`
	LandingLag  uint     `json:"landing_lag"`
	LCancelLag  uint     `json:"l_cancel_lag"`
	AutoCancel  *int     `json:"auto_cancel"`
}

type Special struct {
	ID                 uint     `json:"id" gorm:"primaryKey"`
	Name               string   `json:"name"`
	Start              uint     `json:"start"`
	End                uint     `json:"end"`
	TotalFrames        uint     `json:"total_frames"`
	ShieldStun         *uint    `json:"shield_stun"`
	Damages            []Damage `json:"damages"`
	LandingLag         *uint    `json:"landing_lag"`
	LandingFallSpecial *uint    `json:"landing_fall_special"`
}

type Grab struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Start       uint   `json:"start"`
	TotalFrames uint   `json:"total_frames"`
}

type Throw struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Start       uint     `json:"start"`
	End         uint     `json:"end"`
	TotalFrames uint     `json:"total_frames"`
	Damages     []Damage `json:"damages"`
}

type Dodge struct {
	ID                 uint   `json:"id" gorm:"primaryKey"`
	Name               string `json:"name"`
	Start              uint   `json:"start"`
	End                uint   `json:"end"`
	TotalFrames        uint   `json:"total_frames"`
	LandingFallSpecial *uint  `json:"landing_fall_special"`
}

type Damage struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	MoveID int  `json:"move_id"`
	Number int  `json:"number"`
}

package types

type Song struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title" gorm:"serializer:json"`
	Origin      string `json:"origin" gorm:"serializer:json"`
	PlaysDuring string `json:"plays_during" gorm:"serializer:json"`
}

package entities

type Prices struct {
	SingleMessage int `gorm:"Column:single" json:"single"`
	GroupMessage  int `gorm:"Column:group" json:"group"`
}

package entities

type BlacklistWord struct {
	Word string `gorm:"primaryKey;unique" json:"word"`
}

type BlacklistRegex struct {
	Expression string `gorm:"primaryKey;unique" json:"expression"`
}

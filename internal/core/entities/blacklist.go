package entities

type BlacklistWord struct {
	DBModel
	Word string `gorm:"type:varchar(255);unique" json:"word"`
}

type BlacklistRegex struct {
	DBModel
	Regex string `gorm:"type:string;unique" json:"regex"`
}

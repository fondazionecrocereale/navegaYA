package model

type Clips struct {
	Id       int    `gorm:"type:int;primary_key"`
	Name     string `gorm:"type:varchar(255)"`
	Url      string `gorm:"type:varchar(255)"`
}

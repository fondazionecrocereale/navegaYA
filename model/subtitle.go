package model

import (
	"time"
)

type Subtitle struct {
    Video   int    `gorm:"type:int;primary_key"`
	Index int `gorm:"type:int;primary_key"`
	Start time.Duration
	End   time.Duration
	Text  string `gorm:"type:varchar(255)"`
}

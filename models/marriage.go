package models

import (
	"time"

	"gorm.io/gorm"
)

type Wedding struct {
	gorm.Model
	GroomID     uint
	Groom       People
	BrideID     uint
	Bride       People
	WeddingDate time.Time
}

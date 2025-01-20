package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type bloodType string

const (
	A  bloodType = "A"
	B  bloodType = "B"
	AB bloodType = "AB"
	O  bloodType = "O"
)

type gender string

const (
	P gender = "Pria"
	W gender = "Wanita"
)

type People struct {
	gorm.Model
	Name        string
	IsActive    bool
	Address     *string
	BloodType   *bloodType
	Gender      gender
	PhoneNumber *uint
	Description *string
	BirthPlace  *string
	BirthDate   time.Time
}

func (bt *bloodType) Scan(value interface{}) error {
	*bt = bloodType(value.([]byte))
	return nil
}

func (bt bloodType) Value() (driver.Value, error) {
	return string(bt), nil
}

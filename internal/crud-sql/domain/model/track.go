package model

import "gorm.io/gorm"

type Music struct {
	gorm.Model
	BaseModel

	ID     uint64 `gorm:"primaryKey,autoIncrement"`
	Name   string
	Album  string
	Artist string // create reference to other object
}

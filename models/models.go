package models

import "gorm.io/gorm"

type LogInOut struct {
	gorm.Model
	Password	string	`json:"password"`
	Steps		uint	`json:"steps"`
}
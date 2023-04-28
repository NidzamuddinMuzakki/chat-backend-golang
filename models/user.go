package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
}

type Chat struct {
	gorm.Model
	Name    string `json:"name"`
	Message string `json:"chat"`
}

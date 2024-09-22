package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id          int64
	Title       string
	Author      string
	PublishDate string
}

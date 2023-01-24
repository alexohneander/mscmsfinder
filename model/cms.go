package model

import "gorm.io/gorm"

// Product struct
type CMS struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	UrlPath     string `gorm:"not null" json:"urlpath"`
	QueryString string `gorm:"not null" json:"querystring"`
}

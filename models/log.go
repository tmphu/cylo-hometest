package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Method    string
	URI       string
	Referer   string
	IPAddress string
	Status    int
	UserAgent string
}

package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Category      string
	Price         uint
	FeaturedImage string
	Rating        float32
	Gender        string
	Brand         string
	Color         string
}

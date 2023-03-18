package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Price         uint    `json:"price"`
	FeaturedImage string  `json:"featured_image"`
	Rating        float32 `json:"rating"`
	Gender        string  `json:"gender"`
	Brand         string  `json:"brand"`
	Color         string  `json:"color"`
}

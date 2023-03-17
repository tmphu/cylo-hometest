package main

import (
	"github.com/tmphu/ecom/initializers"
	"github.com/tmphu/ecom/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Log{})
}

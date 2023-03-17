package main

import (
	"github.com/gin-gonic/gin"
	productControllers "github.com/tmphu/ecom/controllers"
	"github.com/tmphu/ecom/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/api/products/get", productControllers.GetProducts)
	r.GET("/api/products/getone", productControllers.GetSingleProduct)
	r.POST("/api/products/create", productControllers.CreateProduct)

	r.Run()
}

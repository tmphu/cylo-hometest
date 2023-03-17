package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/tmphu/ecom/controllers"
	"github.com/tmphu/ecom/initializers"
	"github.com/tmphu/ecom/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.RequestLogger())

	r.GET("/api/products/get", controllers.GetProducts)
	r.GET("/api/products/getone", controllers.GetSingleProduct)
	r.POST("/api/products/create", controllers.CreateProduct)

	r.Run()
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmphu/ecom/services"
)

func GetProducts(c *gin.Context) {
	search := c.Query("search")
	filterGender := c.Query("gender")
	filterCategory := c.Query("category")
	sortPrice := c.Query("sort")

	// create params object
	params := services.ProductQueryParams{
		Search:         search,
		FilterGender:   filterGender,
		FilterCategory: filterCategory,
		SortPrice:      sortPrice,
	}

	// call service
	productService := &services.ProductService{}
	products, err := productService.GetProducts(&params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       products,
	})
}

func GetSingleProduct(c *gin.Context) {
	id := c.Query("id")

	// call service
	productService := &services.ProductService{}
	product, isExist, err := productService.GetSingleProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Internal server error",
		})
		return
	}

	if !isExist {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Product not found",
			"data":       nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       product,
	})
}

func CreateProduct(c *gin.Context) {
	// bind request's body to data struct
	data := services.Data{}
	c.Bind(&data)

	// call service
	productService := &services.ProductService{}
	product, err := productService.CreateProduct(&data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"statusCode": http.StatusCreated,
		"message":    "Success",
		"data":       product,
	})
}

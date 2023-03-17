package productControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmphu/ecom/models"
	"github.com/tmphu/ecom/services"
)

func GetProducts(c *gin.Context) {
	var products []models.Product

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
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": products,
	})
}

func GetSingleProduct(c *gin.Context) {
	var product models.Product
	id := c.Query("id")

	// call service
	productService := &services.ProductService{}
	product, err := productService.GetSingleProduct(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": product,
	})
}

func CreateProduct(c *gin.Context) {
	// bind request's body to data struct
	data := &services.Data{}
	c.Bind(data)

	// call service
	productService := &services.ProductService{}
	product, err := productService.CreateProduct(data)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": product,
	})
}

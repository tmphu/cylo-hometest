package services

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tmphu/ecom/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func createProduct(t *testing.T) models.Product {
	arg := Data{
		Name:          "Test product",
		Category:      "Test category",
		Price:         uint(rand.Intn(100) * 100),
		FeaturedImage: "https://commons.wikimedia.org/wiki/File:Test-Logo.svg",
		Rating:        4.5,
		Gender:        "Male",
		Brand:         "Test brand",
		Color:         "Test color",
	}
	testProductService := ProductService{}
	product, err := testProductService.CreateProduct(&arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Category, product.Category)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.FeaturedImage, product.FeaturedImage)
	require.Equal(t, arg.Rating, product.Rating)
	require.Equal(t, arg.Gender, product.Gender)
	require.Equal(t, arg.Brand, product.Brand)
	require.Equal(t, arg.Color, product.Color)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	createProduct(t)
}

func TestGetSingleProduct(t *testing.T) {
	product1 := createProduct(t)
	testProductService := ProductService{}
	product2, isExist, err := testProductService.GetSingleProduct(strconv.Itoa(int(product1.ID)))
	require.NoError(t, err)
	require.True(t, isExist)
	require.NotEmpty(t, product2)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Category, product2.Category)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.FeaturedImage, product2.FeaturedImage)
	require.Equal(t, product1.Rating, product2.Rating)
	require.Equal(t, product1.Gender, product2.Gender)
	require.Equal(t, product1.Brand, product2.Brand)
	require.Equal(t, product1.Color, product2.Color)
	require.Equal(t, product1.ID, product2.ID)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

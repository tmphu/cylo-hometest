package services

import (
	"fmt"
	"strings"

	"github.com/tmphu/ecom/initializers"
	"github.com/tmphu/ecom/models"
	"gorm.io/gorm"
)

type ProductQueryParams struct {
	Search         string
	FilterGender   string
	FilterCategory string
	SortPrice      string
}

type Data struct {
	Name          string
	Category      string
	Price         uint
	FeaturedImage string
	Rating        float32
	Gender        string
	Brand         string
	Color         string
}

type ProductService struct{}

func (p *ProductService) GetProducts(params *ProductQueryParams) ([]models.Product, error) {
	products := []models.Product{}
	query := initializers.DB

	// search by name
	if s := params.Search; s != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", s))
	}

	// filter by gender
	if fg := params.FilterGender; fg != "" {
		query = query.Where("UPPER(gender) = ?", strings.ToUpper(fg))
	}

	// filter by category
	if fc := params.FilterCategory; fc != "" {
		query = query.Where("UPPER(category) = ?", strings.ToUpper(fc))
	}

	// sort by price
	if sp := params.SortPrice; sp == "price_asc" {
		query = query.Order("price asc")
	} else if sp == "price_desc" {
		query = query.Order("price desc")
	}

	result := query.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (p *ProductService) GetSingleProduct(id string) (models.Product, bool, error) {
	product := models.Product{}
	result := initializers.DB.First(&product, id)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return product, false, result.Error
	}

	if result.Error == gorm.ErrRecordNotFound {
		return product, false, nil
	}
	return product, true, nil
}

func (p *ProductService) CreateProduct(data *Data) (models.Product, error) {
	product := models.Product{Name: data.Name, Category: data.Category, Price: data.Price, FeaturedImage: data.FeaturedImage, Rating: data.Rating, Gender: data.Gender, Brand: data.Brand, Color: data.Color}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

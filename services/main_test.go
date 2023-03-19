package services

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tmphu/ecom/initializers"
)

func init() {
	LoadEnvVariables()
	initializers.ConnectDatabase()
}

func LoadEnvVariables() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

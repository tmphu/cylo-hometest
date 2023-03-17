package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tmphu/ecom/initializers"
	models "github.com/tmphu/ecom/models"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		logEntry := models.Log{
			Method:    c.Request.Method,
			URI:       c.Request.RequestURI,
			Referer:   c.Request.Referer(),
			IPAddress: c.ClientIP(),
			Status:    c.Writer.Status(),
			UserAgent: c.Request.UserAgent(),
		}

		initializers.DB.Create(&logEntry)
	}
}

package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, "No se pudo encontrar un token")
			c.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, "Token inválido")
			c.Abort()
			return
		}
		c.Next()
	}
}
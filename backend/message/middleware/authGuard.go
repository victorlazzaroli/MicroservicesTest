package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlazzaroli/microservicesTest/message/utils"
)

func AuthGuard(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	claims, valid := utils.IsJWTValid(cookie, "parolasegreta")

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	c.Set("user", claims["userData"])

	c.Next()
}

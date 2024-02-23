package middleware

import (
	"fmt"
	"strings"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Unauthorized"})
		return
	}
	token := strings.Split(tokenString, " ")
	parsedToken, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Unauthorized"})
		}

		return []byte("MYSECRETKEY"), nil
	})
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Unauthorized"})
		return
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Unauthorized"})
		return
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Unauthorized"})
		return
	}
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	user, err := models.GetUserById(userId)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Unauthorized"})
		return
	}

	c.Set("user", user)

	c.Next()

}

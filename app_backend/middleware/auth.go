package middleware

import (
	"app_backend/services"
	// "fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// This will help in handling error response
func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

// Authenticate is a middleware that fetches user details from token
func Authenticate(requiredToken string) (string, error) {

	// Get email from encoded token
	splitToken := strings.Split(requiredToken, "bearer ")
	// fmt.Println("Splitoken", splitToken)
	reqToken := splitToken[1]

	userID, err := services.DecodeToken(reqToken)
	return userID, err

}

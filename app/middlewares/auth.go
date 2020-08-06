package middlewares

import (
	"github.com/bracketcore/go-rest-api-template/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func unauthorizedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("Authorization")
		if jwtToken == "" {

			queryToken := c.Query("token")
			if queryToken == "" {
				unauthorizedResponse(c, "Unauthorized access, please login")
			}

			jwtToken = "Bearer " + queryToken
		}

		// Split token string
		extractedToken := strings.Split(jwtToken, "Bearer ")
		if len(extractedToken) != 2 {
			unauthorizedResponse(c, "Unauthorized access, invalid token")
		}

		// Extract token string
		tokenString := strings.TrimSpace(extractedToken[1])
		user, err := utils.CheckAndVerifyToken(tokenString)

		// If user is not in claims return error
		if err != nil {
			unauthorizedResponse(c, "Unauthorized access, invalid token")
		}

		c.Set("auth", user)
		c.Next()
	}
}

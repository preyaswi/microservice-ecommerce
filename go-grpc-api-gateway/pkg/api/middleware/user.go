package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preyaswi/go-grpc-api-gateway/pkg/helper"
	"github.com/preyaswi/go-grpc-api-gateway/pkg/utils/response"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := helper.GetTokenFromHeader(authHeader)
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		userID, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token ", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_id", userID)
		c.Set("user_email", userEmail)
		c.Next()
	}
}

package middleware

import (
	"net/http"

	"example.com/m/auth"
	"github.com/gin-gonic/gin"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "no valid jwt"})
			c.Abort()
			return
		}
		c.Next()
	}
}

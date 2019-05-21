package middleware
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"../config"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		tknstr := c.GetHeader("token")
		if tknstr == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no token provided"})
			return
		}

		claims := &config.JWT{}

		token, err := jwt.ParseWithClaims(tknstr, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		c.Next()
	}
}

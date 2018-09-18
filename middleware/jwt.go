package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JWTMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		if !strings.Contains(authorization, "Bearer") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := string([]byte(authorization)[7:])

		if token == "null" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		type claims struct {
			jwt.StandardClaims
		}

		at(time.Unix(0, 0), func() {
			tok, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil {
				c.AbortWithError(401, err)
			}
			if claims, ok := tok.Claims.(*claims); ok && tok.Valid {
				c.Set("userId", claims.Issuer)
				fmt.Println("current userId", claims.Issuer)
				c.Next()
			} else {
				fmt.Println(err)
				c.AbortWithError(401, err)
				return
			}
		})
	}
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

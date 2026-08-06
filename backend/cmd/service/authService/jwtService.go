package authService

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secretKey []byte
}

func NewJwtService(secret string) *JwtService {
	return &JwtService{
		secretKey: []byte(secret),
	}
}

func (js *JwtService) GenerateToken(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(js.secretKey)
}
func (js *JwtService) ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return js.secretKey, nil
		},
	)

	if err != nil {
		return nil, nil
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, nil
	}

	return claims, nil
}

func (js *JwtService) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "missing token",
			})
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		claims, err := js.ValidateToken(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "invalid token",
			})
			return
		}

		c.Set("userID", claims.Subject)

		c.Next()
	}
}

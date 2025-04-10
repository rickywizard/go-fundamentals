package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Email  string `json:"email"`
	UserID int64  `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(email string, userId int64) (string, error) {
	claims := CustomClaims{
		Email:  email,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			// Issuer:    os.Getenv("JWT_ISSUER"),
			Subject: fmt.Sprint(userId),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	return token.SignedString([]byte(secret))
}

func VerifyToken(token string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}

func GetClaims(context *gin.Context) (*CustomClaims, error) {
	claimsInterface, exists := context.Get("claims")
	if !exists {
		return nil, errors.New("claims not found")
	}

	claims, ok := claimsInterface.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}

	return claims, nil
}

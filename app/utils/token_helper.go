package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CreateToken for user
func CreateToken(user map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 2.0).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// CheckAndVerifyToken used by user
func CheckAndVerifyToken(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Secret), nil
	})

	// Return if there is an error
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if user, ok := claims["user"]; ok {
			return user.(map[string]interface{}), nil
		}
	}

	return nil, fmt.Errorf("unauthorized access: token is invalid or has expired")
}

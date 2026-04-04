package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	Role string `json:"role"`
}

func CreateToken(jwtSecret string, userID string, role string) (string, error) {
	now := time.Now().UTC()
	exp := now.Add(7 * 24 * time.Hour)
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signed, nil
}

func ParseToken(jwtSecret string, tokenString string) (*Claims, error) {
	var claim Claims
	parsedToken, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if claim.Subject == "" {
		return nil, fmt.Errorf("token missing subject")
	}
	return &claim, nil
}

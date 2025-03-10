package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"token":         time.Hour * 2,
	"refresh_token": time.Hour * 72,
}

var jwtSecret = []byte(GetEnv("APP_SECRET", ""))

func GenerateToken(ctx context.Context, userId int, username string, fullname string, tokenType string, email string, now time.Time) (string, error) {
	// Claim token
	claimToken := ClaimToken{
		UserID:   userId,
		Username: username,
		FullName: fullname,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MapTypeToken[tokenType])),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)

	// Sign token
	resultToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return resultToken, fmt.Errorf("failed to generate token: %v", err)
	}
	return resultToken, nil
}

func ValidateToken(ctx context.Context, token string) (*ClaimToken, error) {
	var (
		claimToken *ClaimToken
		ok         bool
	)

	// Parse token
	jwtToken, err := jwt.ParseWithClaims(token, &ClaimToken{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to validate method jwt: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %v", err)
	}

	// Validate token
	if claimToken, ok = jwtToken.Claims.(*ClaimToken); !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return claimToken, nil
}

package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomUserClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type CustomAdminClaims struct {
	AdminID      uint   `json:"admin_id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	IsSuperAdmin bool   `json:"is_super_admin"`
	jwt.RegisteredClaims
}

func GenerateUserToken(userID uint, username string, secret string, duration time.Duration) (string, error) {
	claims := CustomUserClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseUserToken(tokenString string, secret string) (*CustomUserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomUserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func GenerateAdminToken(adminID uint, username string, isSuperAdmin bool, secret string, duration time.Duration) (string, error) {
	claims := CustomAdminClaims{
		AdminID:      adminID,
		Username:     username,
		Role:         "admin",
		IsSuperAdmin: isSuperAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseAdminToken(tokenString string, secret string) (*CustomAdminClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomAdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomAdminClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaim struct {
	UserID       string `json:"user_id"`
	DepartmentID string `json:"department_id"`
	Role         string `json:"role"`
	Type         string `json:"type"` // "access"
	jwt.RegisteredClaims
}

type RefreshTokenClaim struct {
	UserID       string `json:"user_id"`
	DepartmentID string `json:"department_id"`
	Role         string `json:"role"`
	Type         string `json:"type"` // "refresh"
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string, departmentID string, role, secret string, expireMinute int) (string, error) {
	claims := AccessTokenClaim{
		UserID:       userID,
		DepartmentID: departmentID,
		Role:         role,
		Type:         "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireMinute) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userID string, departmentID string, role, secret string, expireDay int) (string, time.Time, error) {
	expireAt := time.Now().Add(time.Duration(expireDay) * 24 * time.Hour)

	claims := RefreshTokenClaim{
		UserID:       userID,
		DepartmentID: departmentID,
		Role:         role,
		Type:         "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	return signed, expireAt, err
}

func ParseAccessToken(tokenString, secret string) (*AccessTokenClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessTokenClaim)
	if !ok || !token.Valid || claims.Type != "access" {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func ParseRefreshToken(tokenString, secret string) (*RefreshTokenClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshTokenClaim)
	if !ok || !token.Valid || claims.Type != "refresh" {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

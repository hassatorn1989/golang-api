package utils

import (
	"time"
	"todo-fiber/internal/dto"

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

func GenerateAccessToken(dto dto.AccessTokenClaimsRequest) (string, error) {
	claims := AccessTokenClaim{
		UserID:       dto.UserID,
		DepartmentID: dto.DepartmentID,
		Role:         dto.Role,
		Type:         "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(dto.ExpiresMinute) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   dto.UserID,
		},
	}

	// stor

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(dto.Secret))
}

func GenerateRefreshToken(dto dto.RefreshTokenClaimsRequest) (string, time.Time, error) {
	expireAt := time.Now().Add(time.Duration(dto.ExpiresDay) * 24 * time.Hour)

	claims := RefreshTokenClaim{
		UserID:       dto.UserID,
		DepartmentID: dto.DepartmentID,
		Role:         dto.Role,
		Type:         "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   dto.UserID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(dto.Secret))
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

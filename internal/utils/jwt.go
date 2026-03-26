package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

type RefreshTokenClaim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint, email, secret string, expireMinute int) (string, error) {
	claims := AccessTokenClaim{
		UserID: userID,
		Email:  email,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireMinute) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userID uint, email, secret string, expireDay int) (string, time.Time, error) {
	expireAt := time.Now().Add(time.Duration(expireDay) * 24 * time.Hour)

	claims := RefreshTokenClaim{
		UserID: userID,
		Email:  email,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   email,
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

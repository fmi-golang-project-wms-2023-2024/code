package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("test")

const (
	accessTokenExp  = time.Minute * 60
	refreshTokenExp = time.Hour * 24
)

func GenerateTokens(sub, role string) (string, string, error) {
	accessToken, err := signJWT(sub, role, accessTokenExp)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := signJWT(sub, role, refreshTokenExp)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func RefreshTokens(refreshToken string) (string, string, error) {
	claims, err := parseJWT(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("invalid refresh token")
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return "", "", fmt.Errorf("refresh token has expired")
	}

	newAccessToken, err := signJWT(claims["sub"].(string), claims["role"].(string), accessTokenExp)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := signJWT(claims["sub"].(string), claims["role"].(string), refreshTokenExp)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func signJWT(sub, role string, exp time.Duration) (string, error) {
	now := time.Now()
	expirationTime := now.Add(exp)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"exp":  expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid JWT")
}

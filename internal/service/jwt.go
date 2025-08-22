package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"johny-tuna/internal/config"
	"johny-tuna/internal/models"
	"time"
)

func GenerateToken(account *models.User) (string, error) {
	key := []byte(config.Env.JwtSecret)

	accessTokenExpirationTime := time.Now().Add(7 * (24 * time.Hour)).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":  account.Id,
		"exp": accessTokenExpirationTime,
	})

	accessTokenString, err := accessToken.SignedString(key)
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	key := []byte(config.Env.JwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func GetIdFromToken(claims jwt.MapClaims) (int64, error) {
	if idFloat, ok := claims["Id"].(float64); ok {
		return int64(idFloat), nil
	}
	if idFloat, ok := claims["id"].(float64); ok {
		return int64(idFloat), nil
	}
	return 0, errors.New("id not found in token")
}

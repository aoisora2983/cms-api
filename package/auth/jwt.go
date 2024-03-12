package auth

import (
	"cms/config"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func GenerateToken(id int, name string, mail string, icon string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(60 * time.Minute)
	jwtSecret = []byte(config.JWTSecretKey())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       "cms",
		"sub":       strconv.Itoa(id),
		"user_id":   id,
		"user_name": name,
		"user_icon": icon,
		"mail":      mail,
		"exp":       jwt.NewNumericDate(expireTime),
		"iat":       jwt.NewNumericDate(nowTime),
		"nbf":       jwt.NewNumericDate(nowTime),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verify(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		},
	)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims.Subject, nil
	}

	return "", err
}

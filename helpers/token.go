package helpers

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
	"wanderloge/structs"
)

func GenerateToken(userName string, userId int, role string) (string, error) {
	secretKey := os.Getenv("SECRET_TOKEN")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"firstName": userName,
		"userId":    userId,
		"role":      role,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Token is invalid")
	}
}

func DecodeToke(tokenString string) (structs.JwtData, error) {
	jwtObject := structs.JwtData{}
	secretKey := os.Getenv("SECRET_TOKEN")
	
	tokenString = strings.TrimSpace(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return jwtObject, fmt.Errorf("Token parsing error: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return jwtObject, fmt.Errorf("Invalid JWT claims")
	}

	jwtObject = structs.JwtData{
		Id:        int(claims["userId"].(float64)),
		FirstName: claims["firstName"].(string),
		Role:      claims["role"].(string),
		Exp:       time.Unix(int64(claims["exp"].(float64)), 0),
	}

	return jwtObject, nil
}

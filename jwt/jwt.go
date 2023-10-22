package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func CreateJWTToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["UserId"] = userId
	TokenDuration := 2 * time.Hour
	claims["exp"] = time.Now().Add(TokenDuration).Unix()
	PrivateKey := viper.GetString("jwt.privateKey")
	res, err := token.SignedString([]byte(PrivateKey))
	return res, err
}

func GetUserIdInJWTToken(token string) (string, error) {
	PrivateKey := viper.GetString("jwt.privateKey")
	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(PrivateKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := payload.Claims.(jwt.MapClaims); ok && payload.Valid {
		UserId := claims["UserId"].(string)
		if err != nil {
			return "", err
		}
		return UserId, nil
	} else {
		return "", errors.New("Cannot get claims in JWT token")
	}
}

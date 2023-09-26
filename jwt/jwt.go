package jwt

import (
	"errors"
	"time"

	"github.com/Lang0808/GolangLibs/cypher"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func CreateJWTToken(userId int32) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	noisedUserId := cypher.NoiseUserId(userId)
	claims["noisedUserId"] = noisedUserId
	TokenDuration := 2 * time.Hour
	claims["exp"] = time.Now().Add(TokenDuration).Unix()
	PrivateKey := viper.GetString("jwt.privateKey")
	res, err := token.SignedString([]byte(PrivateKey))
	return res, err
}

func GetUserIdInJWTToken(token string) (int32, error) {
	PrivateKey := viper.GetString("jwt.privateKey")
	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(PrivateKey), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := payload.Claims.(jwt.MapClaims); ok && payload.Valid {
		NoisedUserId := int64(claims["noisedUserId"].(float64))
		UserId, err := cypher.DenoiseUserId(NoisedUserId)
		if err != nil {
			return -1, err
		}
		return UserId, nil
	} else {
		return 0, errors.New("Cannot get claims in JWT token")
	}
}

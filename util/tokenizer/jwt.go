package tokenizer

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWTToken(payload map[string]interface{}) (string, error) {
	claim := jwt.MapClaims{}
	for k, v := range payload {
		claim[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(viper.GetString("JWT_SECRET")))
}

func VerifyAndParseJWTToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result := map[string]interface{}{}
		for k, v := range claims {
			result[k] = v
		}
		return result, nil
	}

	return nil, err
}

package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
)

var jwtSecret = func() string {
	if os.Getenv("JWT_SECRET") != "" {
		return os.Getenv("JWT_SECRET")
	}
	return "4c73fc2d42887a9ba1678384b611900254a3eeee2a2eeeba57a11eb17c45c47d"
}()

func SignJwt(userId string) (*string, error) {

	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id": userId,
	})

	token, err := sign.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func ParseJwt(token string) (jwt.MapClaims, error) {
	data, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid token")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := data.Claims.(jwt.MapClaims)
	if !ok || !data.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

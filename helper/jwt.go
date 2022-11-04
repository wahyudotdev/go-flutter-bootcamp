package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go-flutter-bootcamp/models/failure"
	"os"
	"time"
)

var jwtSecret = func() string {
	if os.Getenv("JWT_SECRET") != "" {
		return os.Getenv("JWT_SECRET")
	}
	return "4c73fc2d42887a9ba1678384b611900254a3eeee2a2eeeba57a11eb17c45c47d"
}()

func SignJwt(userId string, expired int) (*string, error) {
	mapClaims := jwt.MapClaims{
		"id": userId,
	}
	if expired > 0 {
		minute := time.Minute * time.Duration(expired)
		mapClaims = jwt.MapClaims{
			"id":      userId,
			"expired": time.Now().Add(minute).UnixMilli(),
		}
	}
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), mapClaims)

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
	expired := claims["expired"]
	if expired != nil {
		d1 := expired.(float64)
		d2 := float64(time.Now().UnixMilli())
		if d2 > d1 {
			return claims, errors.New(failure.ExpiredToken)
		}
	}

	return claims, nil
}

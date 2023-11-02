package authentication

import (
	"errors"
	"pharm-stock/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func GenerateToken(id string, username string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = id
	claims["username"] = username
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cfg := configs.Config{}
	return token.SignedString([]byte(cfg.Secret))
}

func ExtractToken(token *jwt.Token) (map[string]any, error) {
	if token.Valid {
		var claims = token.Claims

		expTime, _ := claims.GetExpirationTime()
		if expTime.After(time.Now()) {
			var MapClaim = claims.(jwt.MapClaims)
			newMap := map[string]any{}
			newMap["id"] = MapClaim["id"]
			newMap["username"] = MapClaim["username"]
			newMap["role"] = MapClaim["role"]
			return newMap, nil
		}

		return nil, errors.New("JWT token expired")
	}
	return nil, nil
}

func Middleware() echo.MiddlewareFunc {
	cfg := configs.Config{}
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(cfg.Secret),
		SigningMethod: "HS256",
	})
}
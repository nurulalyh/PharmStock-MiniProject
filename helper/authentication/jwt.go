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

func ExtractTokenUserId(e echo.Context) (string, string, string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		userName := claims["user_name"].(string)
		role := claims["role"].(string)
		return userId, userName, role
	}
	return "", "", ""
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

// func generateToken(signKey string, id string, username string, role string) (string, error) {
// 	var claims = jwt.MapClaims{}
// 	claims["id"] = id
// 	claims["username"] = username
// 	claims["role"] = role
// 	claims["iat"] = time.Now().Unix()
// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

// 	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	validToken, err := sign.SignedString([]byte(signKey))
// 	if err != nil {
// 		return "", errors.New("JWT claims isn't valid, " + err.Error())
// 	}

// 	return validToken, nil
// }

// func generateRefreshToken(signKey string, accessToken string) (string, error) {
// 	var claims = jwt.MapClaims{}
// 	claims["user"] = accessToken
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	refreshToken, err := sign.SignedString([]byte(signKey))
// 	if err != nil {
// 		return "", errors.New("JWT claims isn't valid, " + err.Error())
// 	}

// 	return refreshToken, nil
// }

// func GenerateJWT(signKey string, refreshKey string, userId string, username string, role string) (map[string]any, error) {
// 	var res = map[string]any{}

// 	var accessToken, errGenerateToken = generateToken(signKey, userId, username, role)
// 	if accessToken == "" {
// 		return nil, errors.New("Cannot generate JWT token, " + errGenerateToken.Error())
// 	}

// 	var refreshToken, errRefToken = generateRefreshToken(refreshKey, accessToken)
// 	if refreshToken == "" {
// 		return nil, errors.New("Cannot generate JWT refresh token, " + errRefToken.Error())
// 	}

// 	res["access_token"] = accessToken
// 	res["refresh_token"] = refreshToken

// 	return res, nil
// }

// func RefreshJWT(accessToken string, refreshToken *jwt.Token, signKey string) (map[string]any, error) {
// 	var res = map[string]any{}

// 	expTime, err := refreshToken.Claims.GetExpirationTime()
// 	if err != nil {
// 		return nil, errors.New("Error get token expiration, " + err.Error())
// 	}

// 	if refreshToken.Valid && expTime.After(time.Now()) {
// 		var newClaim = jwt.MapClaims{}

// 		newToken, err := jwt.ParseWithClaims(accessToken, newClaim, func(t *jwt.Token) (interface{}, error) {
// 			return []byte(signKey), nil
// 		})
// 		if err != nil {
// 			return nil, errors.New("JWT error : " + err.Error())
// 		}

// 		newClaim = newToken.Claims.(jwt.MapClaims)
// 		newClaim["iat"] = time.Now().Unix()
// 		newClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

// 		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
// 		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
// 		newSignedRefreshToken, _ := newRefreshToken.SignedString(refreshToken.Signature)

// 		res["access_token"] = newToken.Raw
// 		res["refresh_token"] = newSignedRefreshToken

// 		return res, nil
// 	}

// 	return nil, nil
// }

// func ExtractTokenUserId(e echo.Context) (string, string, string) {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		userId := claims["userId"].(string)
// 		userName := claims["user_name"].(string)
// 		role := claims["role"].(string)
// 		return userId, userName, role
// 	}
// 	return "", "", ""
// }

// func ExtractToken(token *jwt.Token) (any, error) {
// 	if token.Valid {
// 		var claims = token.Claims

// 		expTime, _ := claims.GetExpirationTime()
// 		if expTime.After(time.Now()) {
// 			var MapClaim = claims.(jwt.MapClaims)
// 			newMap := map[string]any{}
// 			newMap["id"] = MapClaim["id"]
// 			newMap["username"] = MapClaim["username"]
// 			newMap["role"] = MapClaim["role"]
// 			return newMap, nil
// 		}

// 		return nil, errors.New("JWT token expired")
// 	}

// 	return nil, nil
// }

// func Middleware() echo.MiddlewareFunc {
// 	cfg := configs.Config{}
// 	return echojwt.WithConfig(echojwt.Config{
// 		SigningKey:    []byte(cfg.Secret),
// 		SigningMethod: "HS256",
// 	})
// }

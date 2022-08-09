package infrastructure

import (
	"errors"
	//"os"
	"fmt"
	"time"

	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	jwt "github.com/dgrijalva/jwt-go"
)

type TokenHandler struct {
	Method *jwt.SigningMethodHMAC
	Secret string
}

func NewTokenHandler() token.TokenHandler {
	tokenHandler := new(TokenHandler)
	tokenHandler.Method = jwt.SigningMethodHS256
	tokenHandler.Secret = "secret" //os.Getenv("CAAG_SECRET")
	if tokenHandler.Secret == "" {
		panic("CAAG_SECRET env variable is not defined.")
	}
	return tokenHandler
}

func (handler *TokenHandler) Generate(uid int, username string, email string) (string, error) {
	// set header
	token := jwt.New(handler.Method)

	// set claims (json contents)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uid
	claims["name"] = username
	claims["email"] = email
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	fmt.Printf("トークンの中身 : %v", claims)

	// signature
	tokenString, err := token.SignedString([]byte(handler.Secret))
	if err != nil {
		return tokenString, errors.New("Failed to generate a new token. " + err.Error())
	}
	return tokenString, nil
}

func (handler *TokenHandler) VerifyToken(tokenString string) error {
	fmt.Println("JWTトークンの中身")
	fmt.Println(tokenString)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.Secret), nil
	})
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return errors.New("Couldn't handle this token:" + err.Error())
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("Not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return errors.New("Token is either expired or not active yet")
		} else {
			return errors.New("Couldn't handle this token:" + err.Error())
		}
	}
	return nil
}

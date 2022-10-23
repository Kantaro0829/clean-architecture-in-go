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

type UserJwt struct {
	Sub   int
	Name  string
	Email string
	Iat   time.Time
	Exp   int64
	jwt.StandardClaims
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
	// // set header
	// token := jwt.New(handler.Method)

	// // set claims (json contents)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = uid
	// claims["name"] = username
	// claims["email"] = email
	// claims["iat"] = time.Now()
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &UserJwt{
		Sub:   uid,
		Email: email,
		Name:  username,
		Iat:   time.Now(),
		Exp:   time.Now().Add(time.Hour * 1).Unix(),
	})

	//fmt.Printf("トークンの中身 : %v", claims)

	// signature
	tokenString, err := token.SignedString([]byte(handler.Secret))
	if err != nil {
		return tokenString, errors.New("Failed to generate a new token. " + err.Error())
	}
	return tokenString, nil
}

//戻り値error から(error, int)に変更
func (handler *TokenHandler) VerifyToken(tokenString string) (int, error) {
	fmt.Println("JWTトークンの中身")
	fmt.Println(tokenString)
	v, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.Secret), nil
	})
	fmt.Println(v)

	userJwt := UserJwt{}
	_, err = jwt.ParseWithClaims(tokenString, &userJwt, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.Secret), nil
	})

	fmt.Println(userJwt.Sub)
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return 0, errors.New("Couldn't handle this token:" + err.Error())
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return 0, errors.New("Not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return 0, errors.New("Token is either expired or not active yet")
		} else {
			return 0, errors.New("Couldn't handle this token:" + err.Error())
		}
	}
	return userJwt.Sub, nil
}

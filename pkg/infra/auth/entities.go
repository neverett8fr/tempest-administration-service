package autho

import "github.com/golang-jwt/jwt/v4"

// https://blog.logrocket.com/jwt-authentication-go/
// https://github.com/golang-jwt/jwt
// https://www.iana.org/assignments/jwt/jwt.xhtml
// https://jwt.io/introduction

// func temp() {
// 	token := jwt.New(jwt.SigningMethodES256)

// }

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

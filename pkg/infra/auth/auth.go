package autho

import (
	"fmt"
	"tempest-administration-service/pkg/infra/db"

	"github.com/golang-jwt/jwt/v4"
)

// https://stackoverflow.com/questions/48855122/keycloak-adaptor-for-golang-application

type TokenProvider struct {
	HMACSigningKey []byte
	Cache          TokenCache
	DB             *db.DBConn
}

func (t *TokenProvider) NewToken(username string, password string) (string, error) {

	// check user database
	// store user passwords here - the other one just for "data"

	err := t.DB.VerifyPassword(username, password)
	if err != nil {
		return "", fmt.Errorf("error verifying password, err %v", err)
	}

	tok := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := tok.SignedString(t.HMACSigningKey)
	if err != nil {
		return "", fmt.Errorf("error signing token, err %v", err)
	}

	t.Cache.CacheToken(tokenString)

	return tokenString, nil

}

func (t *TokenProvider) CheckToken(tok string) error {
	return t.Cache.Exists(tok)

}

func InitialiseTokenProvider(signingKey string, db *db.DBConn) TokenProvider {
	return TokenProvider{
		HMACSigningKey: []byte(signingKey),
		Cache:          NewTokenCache(),
		DB:             db,
	}

}

package autho

import (
	"fmt"
	"tempest-administration-service/pkg/infra/db"

	"github.com/golang-jwt/jwt/v4"
)

// https://stackoverflow.com/questions/48855122/keycloak-adaptor-for-golang-application

type TokenProvider struct {
	Cache TokenCache
	DB    *db.DBConn
}

func (t *TokenProvider) NewToken(username string, password string) (*jwt.Token, error) {

	// check user database
	// store user passwords here - the other one just for "data"

	err := t.DB.VerifyPassword(username, password)
	if err != nil {
		return nil, fmt.Errorf("error verifying password, err %v", err)
	}

	tok := jwt.New(jwt.SigningMethodES256)
	t.Cache.CacheToken(tok.Raw)

	return tok, nil

}

func (t *TokenProvider) CheckToken(tok string) error {
	return t.Cache.Exists(tok)

}

func InitialiseTokenProvider(db *db.DBConn) TokenProvider {
	return TokenProvider{
		Cache: NewTokenCache(),
		DB:    db,
	}

}

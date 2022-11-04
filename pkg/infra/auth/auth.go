package autho

import "github.com/golang-jwt/jwt/v4"

// https://stackoverflow.com/questions/48855122/keycloak-adaptor-for-golang-application

type TokenProvider struct {
	Cache TokenCache
}

func (t *TokenProvider) NewToken() jwt.Token {

	// check user database
	// store user passwords here - the other one just for "data"

	tok := jwt.New(jwt.SigningMethodES256)
	t.Cache.CacheToken(tok.Raw)

	return *tok

}

func InitialiseTokenProvider() TokenProvider {
	return TokenProvider{
		Cache: NewTokenCache(),
	}

}

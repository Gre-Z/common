package jwt

import "github.com/dgrijalva/jwt-go"

func Generate(this jwt.Claims, sercet []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, this)
	return token.SignedString(sercet)
}

func Analysis(this jwt.Claims, token string, sercet []byte) (jwt.Claims, error) {
	if withClaims, err := jwt.ParseWithClaims(token, this, func(token *jwt.Token) (i interface{}, e error) {
		return sercet, nil
	}); err != nil {
		return nil, err
	} else {
		return withClaims.Claims, nil
	}
}

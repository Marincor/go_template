package jwt

import (
	"errors"
	"time"

	"api.default.marincor/config/constants"
	"api.default.marincor/pkg/crypt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Headers struct {
	Key   string
	Value string
}

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

var DefaultHeaders = map[string]string{
	"alg": "RS512",
	"typ": "JWT",
}

func SetupClaims(userEmail string, customArgs ...Headers) jwt.MapClaims {
	accessTokenClaims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(constants.AccessTokenExpirationTime) * time.Minute).Unix(),
		"aud": constants.Audience,
		"jti": uuid.New().String(),
		"iss": userEmail,
		"sub": userEmail,
	}

	if len(customArgs) > 0 {
		for _, arg := range customArgs {
			accessTokenClaims[arg.Key] = arg.Value
		}
	}

	return accessTokenClaims
}

func Verify(tokenString string) (bool, error) {
	tokenParsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return &crypt.ParsePrivateKey().PublicKey, nil
	})
	if err != nil {
		return false, err
	}

	if tokenParsed.Valid {
		return true, nil
	}

	return true, nil
}

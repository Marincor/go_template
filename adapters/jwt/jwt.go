package jwt

import (
	"errors"
	"time"

	"api.default.marincor.pt/config/constants"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	defaultHeaders             = map[string]interface{}{
		"alg": "RS512",
		"typ": "JWT",
	}
)

type JWT struct{}

type Headers struct {
	Key   string
	Value string
}

func New() *JWT {
	return &JWT{}
}

func (jwt *JWT) Validate(token string) bool {
	// TODO: IMPLEMENT VALIDATE JWT
	// client, context := iam.New()

	// auth, err := client.ValidateJWT(context, token, crypt.ParsePrivateKeyToString())
	// if err != nil {
	// 	return false
	// }

	// return auth.Jwt != ""
	return true
}

func (jwt *JWT) Generate(uid string, claims map[string]interface{}, headers map[string]interface{}) (string, error) {
	// TODO: IMPLEMENT GENERATE JWT
	// client, context := iam.New()

	// defaultHeadersInUse := defaultHeaders
	// for headerKey, headerValue := range headers {
	// 	defaultHeadersInUse[headerKey] = headerValue
	// }

	// mappedClaims := []Headers{}
	// for claimKey, claimVal := range claims {
	// 	mappedClaims = append(mappedClaims, Headers{Key: claimKey, Value: fmt.Sprintf("%s", claimVal)})
	// }

	// authToken, err := client.GenerateJWT(context, defaultHeadersInUse, setupClaims(uid, mappedClaims...), crypt.ParsePrivateKeyToString())

	// return authToken.Jwt, err
	return "", nil
}

func setupClaims(uid string, customArgs ...Headers) jwt.MapClaims {
	accessTokenClaims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(constants.AccessTokenExpirationTime) * time.Minute).Unix(),
		"aud": constants.Audience,
		"jti": uuid.New().String(),
		"iss": uid,
		"sub": uid,
	}

	for _, arg := range customArgs {
		accessTokenClaims[arg.Key] = arg.Value
	}

	return accessTokenClaims
}

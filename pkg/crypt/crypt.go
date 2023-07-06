package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"

	"api.default.marincor/adapters/logging"
	"api.default.marincor/entity"
	"golang.org/x/crypto/bcrypt"
)

var ErrToParsePrivate = errors.New("error to parse private key")

func ParsePrivateKey() *rsa.PrivateKey {
	priv, err := os.ReadFile("private.pem")
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "could not get private key file",
			Reason:  err.Error(),
		}, "critical", nil)

		panic("could not get private key file.")
	}

	privatePem, _ := pem.Decode(priv)

	privateKey, err := x509.ParsePKCS8PrivateKey(privatePem.Bytes)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "could not parse private key",
			Reason:  err.Error(),
		}, "critical", nil)

		panic("could not parse private key")
	}

	private, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		go logging.Log(&entity.LogDetails{
			Message: "error to assert type rsa.PrivateKey for private.pem",
		}, "critical", nil)

		panic("error to assert type rsa.PrivateKey for private.pem")
	}

	return private
}

func Encrypt(data string) ([]byte, error) {
	private := ParsePrivateKey()
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "could not decode string to encrypt",
			Reason:  err.Error(),
		}, "critical", nil)

		return nil, err
	}

	plainText, err := rsa.EncryptPKCS1v15(rand.Reader, &private.PublicKey, decoded)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "could not encrypt",
			Reason:  err.Error(),
		}, "critical", nil)

		return nil, err
	}

	return plainText, nil
}

func Decrypt(value string) (string, error) {
	private := ParsePrivateKey()
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "could not decode string to decrypt",
			Reason:  err.Error(),
		}, "critical", nil)

		return "", err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, private, decoded)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "could not decrypt",
			Reason:  err.Error(),
		}, "critical", nil)

		return "", err
	}

	return string(plainText), nil
}

func CheckHash(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))

	return err == nil
}

func ParsePrivateKeyToString() string {
	priv, err := os.ReadFile("private.pem")
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "could not get private key file",
			Reason:  err.Error(),
		}, "critical", nil)

		panic(ErrToParsePrivate)
	}

	privateKey := string(priv)

	return privateKey
}

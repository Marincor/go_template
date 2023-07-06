package totp

import (
	"errors"

	"api.default.marincor/clients/iam"
	"api.default.marincor/config/constants"
)

var errInvalidOTPIntegration = errors.New("provider not integrated with api")

func Generate(secret string) (string, error) {
	if constants.DefaultOTPGenerator == "iam" {
		return generateOTPInIAM(secret)
	}

	return "", errInvalidOTPIntegration
}

func generateOTPInIAM(secret string) (string, error) {
	client, context := iam.New()

	totp, err := client.GenerateOTP(context, secret)
	if err != nil {
		return "", err
	}

	return totp.Token, nil
}

func Validate(totp string, secret string) (bool, error) {
	if constants.DefaultOTPGenerator == "iam" {
		return validateOTPInIAM(totp, secret)
	}

	return false, errInvalidOTPIntegration
}

func validateOTPInIAM(totp string, secret string) (bool, error) {
	client, context := iam.New()

	otpToken, err := client.ValidateOTP(context, totp, secret)
	if err != nil {
		return false, err
	}

	return otpToken.IsValid, nil
}

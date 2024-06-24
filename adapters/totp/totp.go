package totp

import (
	"errors"

	"api.default.marincor.pt/config/constants"
)

type TOTP struct{}

var errInvalidOTPIntegration = errors.New("provider not integrated with api")

func New() *TOTP {
	return &TOTP{}
}

func (otp *TOTP) Generate(secret string) (string, error) {
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

func (otp *TOTP) Validate(totp string, secret string) (bool, error) {
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

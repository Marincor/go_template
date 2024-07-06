package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"api.default.marincor.pt/adapters/logging"
	secretClient "api.default.marincor.pt/clients/google/secretmanager"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/pkg/helpers"
	"github.com/spf13/viper"
)

type Config struct {
	DbHost string
	DbName string
}

func New() *Config {
	if constants.UseSecretManager {
		// TODO: CHECK SECRET MANAGER AWSS
		return setupSecretManager()
	}

	return setupLocal()
}

func setupLocal() *Config {
	var config *Config

	_, file, _, _ := runtime.Caller(0) //nolint: dogsled

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(filepath.Join(filepath.Dir(file), "../"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if constants.Environment == constants.Test {
		log.Printf("Using Test Database")
		config.DbHost = os.Getenv("TEST_DATABASE_URL")
	}

	return config
}

func setupSecretManager() *Config {
	var (
		err    error
		config = &Config{}
	)

	secretList := secretClient.New().ListSecrets(constants.GcpProjectID, constants.SecretPrefix)

	secretBytes, err := helpers.MapToBytes(secretList)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:  "error to parse secrets",
			Reason:   err.Error(),
			Response: secretList,
			Severity: string(constants.SeverityCritical),
		})

		panic(err)
	}

	err = helpers.Unmarshal(secretBytes, config)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:  "error to parse secrets",
			Reason:   err.Error(),
			Response: secretBytes,
			Severity: string(constants.SeverityCritical),
		})

		panic(err)
	}

	return config
}

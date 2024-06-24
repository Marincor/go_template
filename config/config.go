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
	Port               string `json:"port"`
	DBString           string `json:"database_url"`
	DBLogMode          int    `json:"db_log_mode,string"`
	GcpProjectID       string `json:"project_id"`
	StorageBucket      string `json:"storage_bucket"`
	StorageBaseFolder  string `json:"storage_base_folder"`
	MailGunDomain      string `json:"mailgun_domain"`
	MailGunKey         string `json:"mailgun_key"`
	EmailSenderAddress string `json:"email_sender_address"`
	EmailSenderLabel   string `json:"email_sender_label"`
}

func New() *Config {
	if constants.UseSecretManager {
		return setupSecretManager()
	}

	return setupLocal()
}

func setupLocal() *Config {
	var config *Config

	_, file, _, _ := runtime.Caller(0) //nolint: dogsled

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
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
		config.DBString = os.Getenv("TEST_DATABASE_URL")
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
		}, "critical", nil)

		panic(err)
	}

	err = helpers.Unmarshal(secretBytes, config)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:  "error to parse secrets",
			Reason:   err.Error(),
			Response: secretBytes,
		}, "critical", nil)

		panic(err)
	}

	return config
}

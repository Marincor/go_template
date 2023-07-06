package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"api.default.marincor/config/constants"
	"api.default.marincor/pkg/helpers"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

type Config struct {
	Port              string `json:"port"`
	DBString          string `json:"database_url"`
	DBLogMode         int    `json:"db_log_mode"`
	Debug             bool   `json:"debug"`
	GcpProjectID      string `json:"project_id"`
	StorageBucket     string `json:"storage_bucket"`
	StorageBaseFolder string `json:"storage_base_folder"`
	Environment       string `json:"environment"`
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

	return config
}

func setupSecretManager() *Config {
	var (
		config     *Config
		secretList = map[string]interface{}{}
	)

	ctx := context.Background()
	secretClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("Error while trying to connect to Google Cloud Secret Manager, exited with error: %v", err)
	}
	defer secretClient.Close()

	filterPrefix := constants.SecretPrefix
	if filterPrefix == "" {
		filterPrefix = "*"
	}
	secrets := secretClient.ListSecrets(ctx, &secretmanagerpb.ListSecretsRequest{
		Parent: fmt.Sprintf("projects/%s", constants.GcpProjectID),
		Filter: fmt.Sprintf("Name: %s", filterPrefix),
	})
	for {
		secret, err := secrets.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			break
		}

		splitSecret := strings.Split(secret.Name, "/")
		secretName := splitSecret[len(splitSecret)-1]

		if constants.SecretPrefix != "" {
			if strings.Contains(secretName, constants.SecretPrefix) {
				secretName = strings.Split(secretName, constants.SecretPrefix)[1]
			}
		}

		secretList[secretName] = accessSecretVersion(fmt.Sprintf("%s/versions/latest", secret.Name))
	}

	config = &Config{}

	err = helpers.Unmarshal(secretToBytes(secretList), config)
	if err != nil {
		panic("error to parse secrets")
	}

	return config
}

func secretToBytes(secretMap map[string]interface{}) []byte {
	byteSecrets, err := helpers.Marshal(secretMap)
	if err != nil {
		panic("Error to unmarshal configs from Google Cloud")
	}

	return byteSecrets
}

func accessSecretVersion(version string) string {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to create secretmanager client: %v", err))
	}
	defer client.Close()

	result, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: version,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to access secret version: %v", err))
	}

	return string(result.Payload.Data)
}

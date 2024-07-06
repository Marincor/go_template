package secretmanager

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"api.default.marincor.pt/adapters/logging"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/iterator"
)

type GCPSecretManager struct{}

func New() *GCPSecretManager {
	return &GCPSecretManager{}
}

func newClient() (context.Context, *secretmanager.Client) {
	ctx := context.Background()

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:  "error while trying to connect to google cloud secret manager",
			Reason:   err.Error(),
			Severity: string(constants.SeverityCritical),
		})

		panic(err)
	}

	return ctx, client
}

func (secretmanager *GCPSecretManager) ListSecrets(parent string, filterPrefix string) map[string]interface{} {
	ctx, client := newClient()
	defer client.Close()

	filter := filterPrefix
	if filter == "" {
		filter = "*"
	}

	data := &secretmanagerpb.ListSecretsRequest{
		Parent: fmt.Sprintf("projects/%s", parent),
		Filter: fmt.Sprintf("Name: %s", filter),
	}

	secretList := map[string]interface{}{}
	secrets := client.ListSecrets(ctx, data)

	for {
		secret, err := secrets.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			logging.Log(&entity.LogDetails{
				Message:  "error to get next secret in google secret manager",
				Reason:   err.Error(),
				Severity: string(constants.SeverityCritical),
			})

			panic(err)
		}

		latestEnabledVersion := secretmanager.getSecretLastAvailableVersion(secret.Name)

		splitSecret := strings.Split(secret.Name, "/")
		secretName := splitSecret[len(splitSecret)-1]

		if filterPrefix != "" && strings.HasPrefix(secretName, filterPrefix) {
			secretName = strings.TrimPrefix(secretName, filterPrefix)
		}

		secretList[secretName] = secretmanager.accessSecretVersion(latestEnabledVersion.Name)
	}

	return secretList
}

func (secretmanager *GCPSecretManager) accessSecretVersion(version string) string {
	ctx, client := newClient()
	defer client.Close()

	result, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: version,
	})
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:  "error to access secret version in google secret manager",
			Reason:   err.Error(),
			Severity: string(constants.SeverityCritical),
		})

		panic(err)
	}

	return string(result.Payload.Data)
}

func (secretmanager *GCPSecretManager) getSecretLastAvailableVersion(secretName string) *secretmanagerpb.SecretVersion {
	ctx, client := newClient()
	defer client.Close()

	versions := client.ListSecretVersions(ctx, &secretmanagerpb.ListSecretVersionsRequest{
		Parent: secretName,
	})

	var latestEnabledVersion *secretmanagerpb.SecretVersion
	for {
		version, err := versions.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			logging.Log(&entity.LogDetails{
				Message:  "error to get secret versions in google secret manager",
				Reason:   err.Error(),
				Severity: string(constants.SeverityCritical),
			})

			panic(err)
		}

		if version.State != secretmanagerpb.SecretVersion_ENABLED {
			continue
		}

		if latestEnabledVersion == nil || version.CreateTime.Seconds > latestEnabledVersion.CreateTime.Seconds {
			latestEnabledVersion = version
		}
	}

	return latestEnabledVersion
}

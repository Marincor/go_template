package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"api.default.marincor/adapters/logging"
	"api.default.marincor/config/constants"
	"api.default.marincor/entity"
	"api.default.marincor/pkg/app"
	"cloud.google.com/go/storage"
)

func New() (context.Context, *storage.Client) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "error to create gcp storage client",
			Reason:  err.Error(),
		}, "error", nil)

		panic(err)
	}

	return ctx, client
}

func SignedURL(object string, srcFolder string) (string, error) {
	_, client := New()
	defer client.Close()

	finalObject := fmt.Sprintf("%s/%s/%s", app.Inst.Config.StorageBaseFolder, srcFolder, object)

	opts := &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().UTC().Add(constants.SignedURLExp * time.Minute),
	}

	url, err := client.Bucket(app.Inst.Config.StorageBucket).SignedURL(finalObject, opts)
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "error to retrieve google storage signed url",
			Reason:  err.Error(),
		}, "error", nil)

		return "", err
	}

	go logging.Log(&entity.LogDetails{
		Message: "successfully generated signed url",
	}, "error", nil)

	return url, nil
}

func Upload(object string, dstFolder string, reader io.Reader, public bool) error {
	ctx, client := New()

	bucket := client.Bucket(app.Inst.Config.StorageBucket)
	blob := bucket.Object(fmt.Sprintf("%s/%s/%s", app.Inst.Config.StorageBaseFolder, dstFolder, object))
	writer := blob.NewWriter(ctx)

	if _, err := io.Copy(writer, reader); err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "error to copy object",
			Reason:  err.Error(),
		}, "error", nil)

		return err
	}

	if err := writer.Close(); err != nil {
		go logging.Log(&entity.LogDetails{
			Message: "error to close writer",
			Reason:  err.Error(),
		}, "error", nil)

		return err
	}

	if public {
		acl := blob.ACL()
		if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			go logging.Log(&entity.LogDetails{
				Message: "error to set gcp acl to object",
				Reason:  err.Error(),
			}, "error", nil)

			return err
		}
	}

	go logging.Log(&entity.LogDetails{
		Message: fmt.Sprintf("Blob %v uploaded", object),
	}, "debug", nil)

	return nil
}

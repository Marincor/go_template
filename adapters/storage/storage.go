package storage

import (
	"io"

	"api.default.marincor/clients/google/storage"
)

func GetFile(object string, srcFolder string) (string, error) {
	return storage.SignedURL(object, srcFolder)
}

func UploadFile(object string, dstFolder string, reader io.Reader, public bool) error {
	return storage.Upload(object, dstFolder, reader, public)
}

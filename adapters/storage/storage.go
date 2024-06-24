package storage

import (
	"io"

	"api.default.marincor.pt/clients/google/storage"
)

type Storage struct{}

func New() *Storage {
	return &Storage{}
}

func (store *Storage) GetFile(object string, srcFolder string) (string, error) {
	return storage.SignedURL(object, srcFolder)
}

func (store *Storage) UploadFile(object string, dstFolder string, reader io.Reader, public bool) error {
	return storage.Upload(object, dstFolder, reader, public)
}

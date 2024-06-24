package helpers

import (
	"bytes"
	"mime/multipart"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func MapToBytes(datamap map[string]interface{}) ([]byte, error) {
	databyte, err := Marshal(datamap)

	return databyte, err
}

func WriteFormData(formData map[string]string) (*bytes.Buffer, string, error) {
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)

	for key, value := range formData {
		formFieldWriter, err := writer.CreateFormField(key)
		if err != nil {
			return nil, "", err
		}

		_, err = formFieldWriter.Write([]byte(value))
		if err != nil {
			return nil, "", err
		}
	}

	contentType := writer.FormDataContentType()

	writer.Close()

	return form, contentType, nil
}

package fetch

import (
	"bytes"
	"io"
	"mime/multipart"
)

type FormFile struct {
	Filename string
	Contents io.Reader
}

// MakeFormBody simplifies the common task of constructing a multipart/form-data request (i.e. a form submission with files attached)
func MakeFormBody(fields map[string][][]byte, fileFields map[string][]FormFile) (string, io.Reader, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for fieldName, items := range fields {
		for _, item := range items {
			part, err := writer.CreateFormField(fieldName)
			if err != nil {
				return "", nil, err
			}
			part.Write(item)
		}
	}

	for fieldName, files := range fileFields {
		for _, file := range files {
			part, err := writer.CreateFormFile(fieldName, file.Filename)
			if err != nil {
				return "", nil, err
			}
			io.Copy(part, file.Contents)
		}
	}
	writer.Close()

	return writer.FormDataContentType(), body, nil
}

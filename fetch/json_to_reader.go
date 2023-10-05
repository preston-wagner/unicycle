package fetch

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/nuvi/unicycle/defaults"
)

type errReader struct {
	err error
}

func (reader errReader) Read(p []byte) (int, error) {
	return 0, reader.err
}

// JsonToReader simplifies marshalling a struct to json for use in http requests (which accept io.Reader instead of []byte)
func JsonToReader(input any) io.Reader {
	data, err := json.Marshal(input)
	if err != nil {
		return errReader{err: err}
	}
	return bytes.NewReader(data)
}

func ReadString(r io.Reader) (string, error) {
	readerBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(readerBytes), nil
}

func ReadJson[OUTPUT_TYPE any](r io.Reader) (OUTPUT_TYPE, error) {
	readerBytes, err := io.ReadAll(r)
	if err != nil {
		return defaults.ZeroValue[OUTPUT_TYPE](), err
	}

	var output OUTPUT_TYPE
	err = json.Unmarshal(readerBytes, &output)
	return output, err
}

func ReadJsonString[OUTPUT_TYPE any](input string) (OUTPUT_TYPE, error) {
	readerBytes := []byte(input)

	var output OUTPUT_TYPE
	err := json.Unmarshal(readerBytes, &output)
	return output, err
}

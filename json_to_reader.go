package unicycle

import (
	"bytes"
	"encoding/json"
	"io"
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

func ReadJson[OUTPUT_TYPE any](r io.Reader) (OUTPUT_TYPE, error) {
	responseBodyBytes, err := io.ReadAll(r)
	if err != nil {
		return ZeroValue[OUTPUT_TYPE](), err
	}

	var output OUTPUT_TYPE
	err = json.Unmarshal(responseBodyBytes, &output)
	return output, err
}

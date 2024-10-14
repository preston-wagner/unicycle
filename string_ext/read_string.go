package string_ext

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	readerBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(readerBytes), nil
}

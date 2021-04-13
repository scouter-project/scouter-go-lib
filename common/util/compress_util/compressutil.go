package compress_util

import (
	"bytes"
	"compress/gzip"
)

func Gzip(src string) ([]byte, error){
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write([]byte(src))
	if err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil,err
	}
	return buf.Bytes(),nil
}

// Package targz contains methods to create and extract tar gz archives.
//
// Usage (discarding potential errors):
//   	targz.Compress("path/to/the/directory/to/compress", "my_archive.tar.gz")
//   	targz.Extract("my_archive.tar.gz", "directory/to/extract/to")
// This creates an archive in ./my_archive.tar.gz with the folder "compress" (last in the path).
// And extracts the folder "compress" to "directory/to/extract/to/". The folder structure is created if it doesn't exist.
package lib

import (
	"bytes"
	"compress/gzip"
	b64 "encoding/base64"
)

func GzipBase64(data string) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(data)); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	sEnc := b64.StdEncoding.EncodeToString(b.Bytes())
	return sEnc, nil
}

package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
	fmt.Println(SHA1Sig(("config.zip")))
}

/// NOTE: goal is to take a filename, and uncompress the sha1sum value inside the filename
/// Exercise: Decompress only if fine name ends with ".gz"
/// Example: >  cat http.log.gz | gunzip | sha1sum
func SHA1Sig(fileName string) (string, error) {
	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file 

	if strings.HasSuffix(fileName, ".gz") {
	// unzip the file but returning an gzip.Reader pointer which implements Reader interface
		reader, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("%q - gzip: %w", fileName, err)
		}

		defer reader.Close()
		r = reader
	}
	// now find the sha1Sum by using Copy to take a reader and institute a writer
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
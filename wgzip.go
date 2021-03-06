package wgzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//Gzip zips the source to target.
//Appends .gz to file.
func Gzip(source, target string) (string, error) {
	reader, err := os.Open(source)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.gz", filename))
	writer, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	writer.Sync()
	return target, err
}

//UnGzip unzips the source to target.
//Uses archive name for file.
func UnGzip(source, target string) (string, error) {
	reader, err := os.Open(source)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return "", err
	}
	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	writer.Sync()
	return target, err
}

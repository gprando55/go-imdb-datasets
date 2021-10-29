package services

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(destination string, url string) (err error) {

	// Create the file
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// unGzip(destination, targetName)

	return nil
}

func UnGzip(source, targetName string) error {
	// fileName := strings.Split(source, ".")[1]
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	// targetName = filepath.Join(targetName, fileName)
	writer, err := os.Create(targetName)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}

func RemoveFile(source string) error {
	os.Remove(source)
	return nil
}

func SaveJsonFile(destination, data string) (err error) {
	out, err := os.Create(destination)

	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

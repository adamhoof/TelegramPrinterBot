package main

import (
	"io"
	"net/http"
	"os"
)

type FileDownloader struct {
}

const saveTo = "Docs/"

func (fileDownloader *FileDownloader) setup() {

}

func (fileDownloader *FileDownloader) download(fileURL string, fileName string) error {

	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(saveTo + fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

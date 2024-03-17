package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func downloadFile(url string) error {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		// @todo put error into const file
		return logger.Errorf("Can't fetch file", err)
	}
	defer resp.Body.Close()

	filename := strings.ReplaceAll(url, "https://", "")
	filename = strings.ReplaceAll(filename, "/", "-")
	filePath := filepath.Join(settings.destDir, filename)

	out, err := os.Create(filePath)
	if err != nil {
		// @todo better error
		return logger.Errorf("Can't create file", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		// @todo fix
		return logger.Errorf("cant write file", err)
	}

	return nil
}

func imagesDownloader(fileUrls *[]string) error {
	ch := make(chan struct{}, settings.batchSize)
	var wg sync.WaitGroup

	logger.Info("Remove old directory if it exists")
	err := os.RemoveAll(settings.destDir)
	if err != nil {
		return logger.Errorf("error while removing directory", err)
	}

	logger.Info("Create new directory " + settings.destDir)
	err = os.MkdirAll(settings.destDir, os.ModePerm)
	if err != nil {
		return logger.Errorf("error while creating directory", err)
	}

	for _, fileUrl := range *fileUrls {
		wg.Add(1)
		ch <- struct{}{}

		go func(path string) {
			defer func() {
				wg.Done()
				fmt.Println("done", path)
			}()
			fmt.Println("start", path)
			err := downloadFile(path)
			if err != nil {
				fmt.Errorf("can't download file")
			}

			<-ch
		}(fileUrl)

	}

	wg.Wait()

	return nil
}

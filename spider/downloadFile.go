package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func downloadFile(url, destDir string) error {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		// @todo put error into const file
		return fmt.Errorf("unable to dl the file %v", err)
	}
	resp.Body.Close()

	// Create the destination directory if it doesn't exist
	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error while creating directory: %v", err)
	}

	filename := strings.ReplaceAll(url, "/", "-")
	filePath := filepath.Join(destDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		// @todo better error
		return fmt.Errorf("Can't create file %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		// @todo fix
		return fmt.Errorf("cant write file %v", err)
	}

	return nil
}

// Second loop that will batch stuff
//func DownloadAllFiles(fileUrl string[]) error {
//	ch := make(chan struct{})
//
//}

// First loop on all download

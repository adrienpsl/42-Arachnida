package main

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetLinksFromHTML(url string) ([]string, error) {
	bodyText, htmlError := getHTMLBody(url)
	if htmlError != nil {
		log.Println("Exiting due to HTML error")
		return nil, htmlError
	}

	re := regexp.MustCompile(`https?://[^\s<>"']*[^.,;'">\s<>()\]\!]`)
	matches := re.FindAllString(bodyText, -1)

	return matches, nil
}

func getHTMLBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(BadUrlError)
		return "", err
	}
	defer closeBody(resp.Body)

	bodyBytes, bodyErr := io.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Println(bodyErr.Error())
		return "", bodyErr
	}

	return string(bodyBytes), nil
}

func closeBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
}

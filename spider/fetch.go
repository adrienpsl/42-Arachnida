package main

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
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

func FilterLinksByExtensions(links []string, extensions []string, mode bool) []string {
	var filteredLinks []string

	for _, link := range links {
		hasMatchingExtension := false
		for _, ext := range extensions {
			if strings.HasSuffix(link, ext) {
				hasMatchingExtension = true
				break
			}
		}

		if (mode && hasMatchingExtension) || (!mode && !hasMatchingExtension) {
			filteredLinks = append(filteredLinks, link)
		}
	}

	return filteredLinks
}

func FilterNewLinks(links []string, oldLinks []string) []string {
	var newLinks []string

	for _, link := range links {
		if !sameString(oldLinks, link) {
			newLinks = append(newLinks, link)
		}
	}

	return newLinks
}

func sameString(links []string, link string) bool {
	for _, l := range links {
		if l == link {
			return true
		}
	}
	return false
}

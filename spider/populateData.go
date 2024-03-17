package main

import (
	"fmt"
	"sync"
)

var visitedLinks []string
var foundedImages []string

func LoopOnLinks(data map[int]Data) map[int]Data {
	var mu sync.Mutex

	visitedLinks = []string{settings.startUrl}
	foundedImages = []string{}
	links := []string{settings.startUrl}
	for index := 0; index < settings.deep; index++ {
		fmt.Println("start loop", links, index, settings.deep)
		data[index] = Data{Images: make([]string, 0), Links: make([]string, 0)}
		data[index] = PopulateData(data[index], links, &mu)
		links = data[index].Links
		visitedLinks = append(visitedLinks, links...)
	}

	return data
}

func PopulateData(data Data, links []string, mu *sync.Mutex) Data {
	ch := make(chan struct{}, settings.batchSize)
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		ch <- struct{}{}

		go func(url string) {
			defer func() {
				fmt.Println("done", url)
				wg.Done()
			}()

			fmt.Println("star", url)
			rawLinks, err := GetLinksFromHTML(url)
			if err != nil {
				fmt.Println("bad url continue")
				return
			}
			// i can have multiple images with same name, i need to keep all and update their name.
			allImages := FilterLinksByExtensions(rawLinks, settings.extensions, true)
			notImageLinks := FilterLinksByExtensions(rawLinks, settings.extensions, false)

			// here i lock to be sure that each element is uniq
			mu.Lock()
			images := FilterNewLinks(allImages, foundedImages)
			foundedImages = append(foundedImages, images...)

			links = FilterNewLinks(notImageLinks, visitedLinks)
			visitedLinks = append(visitedLinks, links...)
			mu.Unlock()

			data.Images = append(data.Images, images...)
			data.Links = append(data.Links, links...)
			<-ch
		}(link)
	}

	wg.Wait()

	return data
}

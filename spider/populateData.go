package main

import (
	"fmt"
	"sync"
)

var extensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}

const batchSize = 30

func PopulateData(data Data, links []string, allReadyVisited *[]string, allReadyHaveImage *[]string, mu *sync.Mutex) Data {
	ch := make(chan struct{}, batchSize)
	var wg sync.WaitGroup

	fmt.Println("links", links)
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
			allImages := FilterLinksByExtensions(rawLinks, extensions, true)
			notImageLinks := FilterLinksByExtensions(rawLinks, extensions, false)

			mu.Lock()
			images := FilterNewLinks(allImages, *allReadyHaveImage)
			*allReadyHaveImage = append(*allReadyHaveImage, images...)

			links = FilterNewLinks(notImageLinks, *allReadyVisited)
			*allReadyVisited = append(*allReadyVisited, links...)
			mu.Unlock()

			data.Images = append(data.Images, images...)
			data.Links = append(data.Links, links...)
			<-ch
		}(link)
	}

	wg.Wait()

	return data
}

func LoopOnLinks(data map[int]Data, startUrl string, deep int) map[int]Data {
	var mu sync.Mutex

	allReadyVisited := []string{startUrl}
	allReadyHaveImage := []string{}
	links := []string{startUrl}
	for index := 0; index < deep; index++ {
		fmt.Println("start loop", links, index, deep)
		data[index] = Data{Images: make([]string, 0), Links: make([]string, 0)}
		data[index] = PopulateData(data[index], links, &allReadyVisited, &allReadyHaveImage, &mu)
		links = data[index].Links
		allReadyVisited = append(allReadyVisited, links...)
	}

	return data
}

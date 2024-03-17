package main

import (
	"fmt"
	"sync"
)

var visitedLinks []string
var mu sync.Mutex

func LoopOnLinks(data *Data) {
	deepData := (*data).DeepData

	visitedLinks = []string{settings.startUrl}
	links := []string{settings.startUrl}
	for index := 0; index < settings.deep; index++ {
		fmt.Println("start loop", links, index, settings.deep)
		deepData[index] = DeepData{Images: make([]string, 0), Links: make([]string, 0)}
		PopulateData(&deepData[index], links, data)
		links = deepData[index].Links
		visitedLinks = append(visitedLinks, links...)
	}

}

func PopulateData(deepData *DeepData, links []string, data *Data) {
	ch := make(chan struct{}, settings.batchSize)
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		ch <- struct{}{}

		go func(url string) {
			defer func() {
				//fmt.Println("done", url)
				wg.Done()
			}()

			//fmt.Println("star", url)
			rawLinks, err := GetLinksFromHTML(url)
			if err != nil {
				//fmt.Println("bad url continue")
				return
			}
			// i can have multiple images with same name, i need to keep all and update their name.
			allImages := FilterLinksByExtensions(rawLinks, settings.extensions, true)
			notImageLinks := FilterLinksByExtensions(rawLinks, settings.extensions, false)

			// here i lock to be sure that each element is uniq
			mu.Lock()
			images := FilterNewLinks(allImages, data.FoundedImages)
			data.FoundedImages = append(data.FoundedImages, images...)

			links = FilterNewLinks(notImageLinks, visitedLinks)
			visitedLinks = append(visitedLinks, links...)
			mu.Unlock()

			deepData.Images = append(deepData.Images, images...)
			deepData.Links = append(deepData.Links, links...)
			<-ch
		}(link)
	}

	wg.Wait()
}

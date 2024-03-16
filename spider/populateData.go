package main

import (
	"fmt"
	"sync"
)

var extensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}

func PopulateData(data Data, links []string, allReadyVisited []string) Data {
	for _, link := range links {
		rawLinks, _ := GetLinksFromHTML(link)
		// i can have multiple images with same name, i need to keep all and update their name.
		images := FilterLinksByExtensions(rawLinks, extensions, true)

		notImageLinks := FilterLinksByExtensions(rawLinks, extensions, false)
		links = FilterNewLinks(notImageLinks, allReadyVisited)

		data.Images = append(data.Images, images...)
		data.Links = append(data.Links, links...)
	}

	return data
}

func LoopOnLinks(data map[int]Data, startUrl string, deep int) map[int]Data {
	var wg sync.WaitGroup
	allReadyVisited := []string{startUrl}
	batchSize := 10

	// Create a channel to control the number of concurrent goroutines
	semaphore := make(chan struct{}, batchSize)

	links := []string{startUrl}
	for i := 0; i < deep; i++ {
		// Acquire a slot in the semaphore
		semaphore <- struct{}{}

		// Increment the WaitGroup counter
		wg.Add(1)

		// Launch a goroutine for each iteration
		go func(index int) {
			defer func() {
				// Signal the completion of the goroutine
				wg.Done()

				// Release the slot in the semaphore
				<-semaphore
				fmt.Println("done")
			}()

			fmt.Println("start loop", links, index, deep)
			data[index] = Data{Images: make([]string, 0), Links: make([]string, 0)}
			data[index] = PopulateData(data[index], links, allReadyVisited)
			links = data[index].Links
			allReadyVisited = append(allReadyVisited, links...)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return data
}

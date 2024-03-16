package main

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
	allReadyVisited := []string{startUrl}

	links := []string{URL}
	for i := 0; i < deep; i++ {
		data[i] = Data{Images: make([]string, 0), Links: make([]string, 0)}
		data[i] = PopulateData(data[i], links, allReadyVisited)
		links = data[i].Links
		allReadyVisited = append(allReadyVisited, links...)
	}
	return data
}

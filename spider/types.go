package main

type DeepData struct {
	Images []string
	Links  []string
}

type Data struct {
	DeepData      []DeepData `json:"deep-data"`
	FoundedImages []string   `json:"founded-images"`
}

type Settings struct {
	batchSize  int
	deep       int
	extensions []string
	startUrl   string
	destDir    string
}

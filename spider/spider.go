package main

import (
	"os"
)

func main() {
	parseArg(&settings)
	if (settings.deep < 0) || (settings.batchSize < 0) {
		logger.Error("deep and batch size must be positive")
		os.Exit(42)
	}

	data := Data{
		DeepData:      make([]DeepData, settings.deep),
		FoundedImages: []string{},
	}

	LoopOnLinks(&data)

	err := saveJsonDebug(data)
	if err != nil {
		logger.Error("Error while saving json debug exiting")
	}

	err = imagesDownloader(&data.FoundedImages)
	if err != nil {
		logger.Error("Error while downloading images exiting")
		os.Exit(42)
	}
}

package main

import (
	"flag"
	"strings"
)

func parseArg(settings *Settings) {
	logger.Info("Parsing arguments")

	recursive := flag.Bool("r", true, "Is the search recursive")
	deep := flag.Int("l", 2, "The depth of the recursive search")
	destDir := flag.String("p", "data", "The destination directory for the downloaded images")
	batchSize := flag.Int("b", 30, "The batch size for the concurrent requests")
	extensions := flag.String("e", ".jpg,.jpeg,.png,.gif,.bmp", "The extensions of the images to download")

	flag.Parse()
	url := flag.Arg(0)

	logger.Info("No recursive option, setting deep to 0")
	if !*recursive {
		*deep = 0
	}

	*settings = Settings{
		batchSize:  *batchSize,
		deep:       *deep,
		destDir:    *destDir,
		extensions: strings.Split(*extensions, ","),
		startUrl:   url,
	}
}

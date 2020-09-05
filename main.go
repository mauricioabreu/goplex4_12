package main

import (
	"flag"

	"github.com/mauricioabreu/goplex4_12/xkcd"
)

func main() {
	var index bool
	flag.BoolVar(&index, "index", false, "Index xkcd comics in a JSON file")
	flag.Parse()

	if index {
		indexer := xkcd.NewIndexer("./data/comics.json", "https://xkcd.com", 10)
		indexer.Index()
	}
}

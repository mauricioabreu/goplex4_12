package main

import (
	"github.com/mauricioabreu/goplex4_12/xkcd"
)

func main() {
	indexer := xkcd.NewIndexer("./data/comics.json", "https://xkcd.com", 10)
	indexer.Index()
}

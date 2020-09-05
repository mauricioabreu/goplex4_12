package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mauricioabreu/goplex4_12/xkcd"
)

func main() {
	var index bool
	var query string
	flag.BoolVar(&index, "index", false, "Index xkcd comics in a JSON file")
	flag.StringVar(&query, "query", "", "Terms to search for comics")
	flag.Parse()

	indexer := xkcd.NewIndexer("./data/comics.json", "https://xkcd.com", 10)

	if index {
		indexer.Index()
	} else if query != "" {
		comics, err := indexer.LoadComics()
		if err != nil {
			log.Fatal(err)
		}
		results := xkcd.Search(query, comics)
		for _, comic := range results {
			fmt.Println(comic)
		}
	}
}

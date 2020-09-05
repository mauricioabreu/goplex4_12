package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/mauricioabreu/goplex4_12/xkcd"
)

var comicTemplate = template.Must(template.New("comic").Parse(`
Number: {{.Num}}
Title: {{.Title}}
Image: {{.Img}}
Transcript: {{.Transcript}}
Link: {{.Link}}
`))

func main() {
	var index bool
	var query string
	var maxComics int
	flag.BoolVar(&index, "index", false, "Index xkcd comics in a JSON file")
	flag.StringVar(&query, "query", "", "Terms to search for comics")
	flag.IntVar(&maxComics, "max_comics", 10, "How many comics you want to index?")
	flag.Parse()

	indexer := xkcd.NewIndexer("./data/comics.json", "https://xkcd.com", maxComics)

	if index {
		indexer.Index()
	} else if query != "" {
		comics, err := indexer.LoadComics()
		if err != nil {
			log.Fatal(err)
		}
		results := xkcd.Search(query, comics)
		for _, comic := range results {
			err = comicTemplate.Execute(os.Stdout, comic)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

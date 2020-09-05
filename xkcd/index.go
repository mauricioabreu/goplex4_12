package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Indexer contains information about comics indexing
type Indexer struct {
	DataPath  string
	BaseURL   string
	MaxComics int
}

// NewIndexer create a new indexer ready to fetch comics
func NewIndexer(dataPath, baseURL string, maxComics int) *Indexer {
	return &Indexer{
		DataPath:  dataPath,
		BaseURL:   baseURL,
		MaxComics: maxComics,
	}
}

// Index index all comics
func (i *Indexer) Index() {
	var comics []*Comic
	for c := 1; c <= i.MaxComics; c++ {
		comic, err := i.Fetch(c)
		fmt.Println(comic)
		if err != nil {
			log.Println(err)
		}
		comics = append(comics, comic)
	}

	comicsJSON, err := json.MarshalIndent(comics, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(i.DataPath, comicsJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Fetch get comic from xkcd site
func (i *Indexer) Fetch(id int) (*Comic, error) {
	url := i.URLForComic(id)
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/vnd.github.v3+json")
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch comic: %s", response.Status)
	}

	var comic Comic
	if err := json.NewDecoder(response.Body).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}

// LoadComics load all comics in memory
func (i *Indexer) LoadComics() ([]*Comic, error) {
	var comics []*Comic
	data, err := ioutil.ReadFile(i.DataPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

// URLForComic build a URL for the comic
func (i *Indexer) URLForComic(id int) string {
	return fmt.Sprintf("%s/%d/info.0.json", i.BaseURL, id)
}

package xkcd

import "strings"

// Search retrieve comics for the given query text
func Search(query string, comics []*Comic) []*Comic {
	var foundComics []*Comic
	for _, c := range comics {
		if contains(query, c) {
			foundComics = append(foundComics, c)
		}
	}
	return foundComics
}

func contains(query string, c *Comic) bool {
	if strings.Contains(c.Title, query) {
		return true
	}

	return false
}

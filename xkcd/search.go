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
	lQuery := strings.ToLower(query)
	if strings.Contains(strings.ToLower(c.Title), lQuery) {
		return true
	}
	if strings.Contains(strings.ToLower(c.SafeTitle), lQuery) {
		return true
	}
	if strings.Contains(strings.ToLower(c.Transcript), lQuery) {
		return true
	}
	if strings.Contains(strings.ToLower(c.Alt), lQuery) {
		return true
	}

	return false
}

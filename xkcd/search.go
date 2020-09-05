package xkcd

import "strings"

func search(text string, comics []*Comic) []*Comic {
	var foundComics []*Comic
	for _, c := range comics {
		if contains(text, c) {
			foundComics = append(foundComics, c)
		}
	}
	return foundComics
}

func contains(text string, c *Comic) bool {
	if strings.Contains(text, c.Title) {
		return true
	}

	return false
}

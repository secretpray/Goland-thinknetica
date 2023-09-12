package index

import (
	"strings"
	"thinknetica/Lesson-12/pkg/crawler"
)

type ReverseIndex struct {
	storage map[string][]int
}

func NewReverseIndex() *ReverseIndex {
	return &ReverseIndex{storage: make(map[string][]int)}
}

func (r *ReverseIndex) Build(data []crawler.Document) map[string][]int {
	result := make(map[string][]int)

	for _, value := range data {
		tokens := tokenize(value.Title)
		for _, token := range tokens {
			if token == "-" || token == "&" || token == "." {
				continue
			}

			result[token] = append(result[token], value.ID)
		}
	}

	return result
}

func tokenize(input string) []string {
	tokens := strings.Fields(input)
	return tokens
}

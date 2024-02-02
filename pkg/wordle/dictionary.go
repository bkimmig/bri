package wordle

import (
	_ "embed"
	"strings"
)

// the embed package allows you to embed files into your Go code
// in this case we load in the words as a string and then parse them
// to get them in the right format.
var (
	letterCount = 5

	//go:embed data/words.txt
	dictionary string
)

func loadWords() ([]string, error) {
	allWords := strings.Split(dictionary, "\n")
	words := make([]string, 0, len(allWords))
	for _, word := range allWords {
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)
		if len(word) != letterCount {
			continue
		}
		words = append(words, word)
	}
	return words, nil
}

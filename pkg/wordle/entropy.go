package wordle

import (
	"math"
)

func GetDictionaryEntropy() (map[string]float64, error) {
	words, err := loadWords()
	if err != nil {
		return nil, err
	}
	entropy := entropy(words)
	return entropy, nil
}

func entropy(words []string) map[string]float64 {
	entropy := make(map[string]float64)
	for _, word := range words {
		entropy[word] = wordEntropy(word)
	}

	return entropy
}

func wordEntropy(word string) float64 {
	numLetters := float64(len(word))
	counts := letterCounts(word)
	var entropy float64
	for _, letter := range word {
		p := counts[string(letter)] / numLetters
		entropy += p * math.Log2(p)
	}
	return -1 * entropy
}

func letterCounts(word string) map[string]float64 {
	counts := make(map[string]float64)
	for _, letter := range word {
		counts[string(letter)] += 1
	}

	return counts
}

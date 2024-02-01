package wordle

import (
	"math"
)

const (
	EntropyV1 = "v1"
	EntropyV2 = "v2"
)

func GetDictionaryEntropy(version string) (map[string]float64, error) {
	words, err := loadWords()
	if err != nil {
		return nil, err
	}
	switch version {
	case EntropyV1:
		return entropy(words), nil
	case EntropyV2:
		return entropyV2(words), nil
	default:
		return entropy(words), nil
	}
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

func entropyV2(words []string) map[string]float64 {
	totalWords := float64(len(words))
	counts := make(map[string]float64)
	for _, word := range words {

		uniqueLetters := make(map[string]struct{}, len(word))
		for _, letter := range word {
			uniqueLetters[string(letter)] = struct{}{}
		}

		for letter := range uniqueLetters {
			counts[letter] += 1
		}
	}

	entropy := make(map[string]float64, len(counts))
	for _, word := range words {
		e := 0.0
		for _, letter := range word {
			c, ok := counts[string(letter)]
			if !ok {
				continue
			}
			p := c / totalWords
			e += p * math.Log2(p)
		}
		entropy[word] = -1 * e
	}

	return entropy
}

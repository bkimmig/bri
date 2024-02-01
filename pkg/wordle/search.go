package wordle

import "fmt"

const (
	miss     = "0"
	hit      = "1"
	exactHit = "2"
)

type criteria struct {
	excludes  map[string]struct{}
	incorrect map[string]map[int]struct{}
	exact     map[int]string
}

func (c criteria) isValid(word string) bool {
	containsLetter := make(map[string]struct{})

	for i, b := range word {
		letter := string(b)

		// contains an excluded letter? word can be removed immediately
		if _, ok := c.excludes[letter]; ok {
			return false
		}

		// do I know that the letter in this word belngs in
		l, ok := c.exact[i]
		if ok {
			if l != letter {
				return false
			}
		}

		pos, ok := c.incorrect[letter]
		if ok {
			_, ok := pos[i]
			if ok {
				return false
			}
			containsLetter[letter] = struct{}{}
		}
	}

	return len(containsLetter) == len(c.incorrect)
}

func (c criteria) Update(word string, info string) error {
	for i, b := range word {
		letter := string(b)
		letterInfo := string(info[i])

		switch letterInfo {
		case miss:
			c.excludes[letter] = struct{}{}
		case hit:
			if _, ok := c.incorrect[letter]; !ok {
				c.incorrect[letter] = make(map[int]struct{})
			}
			c.incorrect[letter][i] = struct{}{}
		case exactHit:
			if _, ok := c.exact[i]; !ok {
				c.exact[i] = letter
			}
		default:
			return fmt.Errorf("invalid info: %s", info)
		}
	}

	return nil
}

func NewCriteria() criteria {
	return criteria{
		excludes:  make(map[string]struct{}),
		incorrect: make(map[string]map[int]struct{}),
		exact:     make(map[int]string),
	}
}

// Search takes the entropy and criteria and returns the subset of the entropy that matches the criteria
// entropy data object is structure like is word: entropy
func Search(entropy map[string]float64, c criteria) map[string]float64 {
	subsetEntropy := make(map[string]float64)
	for word, e := range entropy {
		if c.isValid(word) {
			subsetEntropy[word] = e
		}
	}
	return subsetEntropy
}

package wordle

import "fmt"

const (
	miss     = "0"
	hit      = "1"
	exactHit = "2"
)

type criteria struct {
	info     map[string]bool
	contains map[string]struct{}
	exact    map[string]struct{}
}

func (c criteria) isValid(word string) bool {
	contains := make(map[string]struct{}, len(word))
	containsExact := make(map[string]struct{}, len(word))
	for i, b := range word {
		letter := string(b)
		key := c.key(letter, i)

		// if something is false, throw it out immediately
		info, ok := c.info[key]
		if ok && !info {
			return false
		}

		// if the letter is not in the word, throw it out immediately
		info, ok = c.info[letter]
		if ok && !info {
			return false
		}

		// check if the letter is contained in the word
		_, ok = c.contains[letter]
		if ok {
			contains[letter] = struct{}{}
		}

		// check if the letter is in the correct position
		_, ok = c.exact[key]
		if ok {
			containsExact[key] = struct{}{}
		}
	}

	// number of uniuqe letters contained in the word must match the criteria
	return len(contains) == len(c.contains) && len(containsExact) == len(c.exact)
}

func (c criteria) Update(word, info string) error {
	for i, b := range word {
		letter := string(b)
		letterInfo := string(info[i])
		key := c.key(letter, i)

		switch letterInfo {
		case miss:
			c.info[letter] = false
			c.info[key] = false
		case hit:
			c.info[key] = false
			c.contains[letter] = struct{}{}
		case exactHit:
			c.info[key] = true
			c.contains[letter] = struct{}{}
			c.exact[key] = struct{}{}
		default:
			return fmt.Errorf("invalid info: %s", info)
		}
	}

	return nil
}

func (c criteria) key(letter string, position int) string {
	return fmt.Sprintf("%s-%d", letter, position)
}

func NewCriteria() criteria {
	return criteria{
		info:     make(map[string]bool),
		contains: make(map[string]struct{}),
		exact:    make(map[string]struct{}),
	}
}

// Search takes the entropy and criteria and returns the subset of the entropy that matches the criteria
// entropy data object is structure like is word: entropy
func Search(entropy map[string]float64, c criteria) map[string]float64 {
	subsetEntropy := make(map[string]float64, len(entropy))
	for word, e := range entropy {
		if c.isValid(word) {
			subsetEntropy[word] = e
		}
	}
	return subsetEntropy
}

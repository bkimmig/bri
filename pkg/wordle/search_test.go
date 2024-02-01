package wordle

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/maps"
)

func TestSearch(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		c := NewCriteria()
		entropy := map[string]float64{
			"stray": 1.,
			"store": 1.,
			"ivory": 1.,
			"ebola": 1.,
			"huffy": 1.,
			"giddy": 1.,
			"buzzy": 1.,
			"missy": 1.,
		}

		// actual word is ivory
		c.Update("stray", "10002")
		e := Search(entropy, c)

		fmt.Println(c)

		got := maps.Keys(e)
		sort.Strings(got)
		want := []string{"missy", "giddy", "huffy", "buzzy"}
		sort.Strings(want)

		if !cmp.Equal(got, want) {
			t.Fatalf("diff %s", cmp.Diff(got, want))
		}

	})
}

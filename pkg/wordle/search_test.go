package wordle

import (
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
			"fussy": 1.,
			"biome": 1.,
		}

		// actual word is ivory
		err := c.Update("stray", "10002")
		if err != nil {
			t.Fatal(err)
		}
		e := Search(entropy, c)

		got := maps.Keys(e)
		sort.Strings(got)
		want := []string{"missy", "fussy"}
		sort.Strings(want)

		if !cmp.Equal(got, want) {
			t.Fatalf("diff %s", cmp.Diff(got, want))
		}
	})

	t.Run("search-2", func(t *testing.T) {
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
			"fussy": 1.,
			"biome": 1.,
			"brown": 1.,
			"bourd": 1.,
			"boxer": 1.,
		}

		// actual word is boxer
		updates := [][]string{
			{"stray", "00100"},
			{"brown", "21100"},
		}
		for _, u := range updates {
			err := c.Update(u[0], u[1])
			if err != nil {
				t.Fatal(err)
			}
		}
		e := Search(entropy, c)

		got := maps.Keys(e)
		sort.Strings(got)
		want := []string{"bourd", "boxer"}
		sort.Strings(want)

		if !cmp.Equal(got, want) {
			t.Fatalf("diff %s", cmp.Diff(got, want))
		}
	})
}

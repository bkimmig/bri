package wordle

import (
	"testing"
)

func TestLoadWords(t *testing.T) {

	t.Run("dictionary", func(t *testing.T) {
		words, err := loadWords()
		if err != nil {
			t.Fatal(err)
		}
		want := 14855
		got := len(words)
		if want != got {
			t.Fatalf("got %d words, want %d", got, want)
		}
	})
}

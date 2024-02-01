package wordle

import (
	"testing"
)

func TestLoadDictionary(t *testing.T) {
	t.Run("macos", func(t *testing.T) {
		words, err := loadDictionary(macOSDict)
		if err != nil {
			t.Fatal(err)
		}
		want := 10239
		got := len(words)
		if want != got {
			t.Fatalf("got %d words, want %d", got, want)
		}
	})
}

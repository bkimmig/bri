package wordle

import (
	"fmt"
	"testing"
)

func TestGetDictionaryEntropy(t *testing.T) {
	t.Run("entropy", func(t *testing.T) {
		entropy, err := GetDictionaryEntropy()
		if err != nil {
			t.Fatal(err)
		}

		if len(entropy) == 0 {
			t.Fatal("no entropy found")
		}

		fmt.Println("here", entropy["ebony"])
		t.Fatal("bork")

	})
}

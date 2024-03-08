package cmd

import (
	"bri/pkg/wordle"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

var wordleLong = `A command for helping solve a wordle. It will prompt you for
the word and the exact hits (2), hits (1), and misses (0) for each
letter.

For example, if the wordle word is 'store' and you guessed 'stray', the input would be 'stray 22100'

It will then use the entropy of the dictionary to find the most
likely words.`

func wordleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wordle",
		Short: "A command for helping solve a wordle",
		Long:  wordleLong,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Input your first word and the exact hits (2), hits (1), and misses (0) for each letter.\n")
			fmt.Print("For example, if the wordle word is 'store' and you guessed 'stray', the input would be 'stray 22100'.\n")

			criteria := wordle.NewCriteria()
			entropy, err := wordle.GetDictionaryEntropy(wordle.EntropyV2)
			if err != nil {
				fmt.Printf("error getting entropy: %v\n", err)
			}

			// set up the readline for prompt input
			l, err := readline.NewEx(&readline.Config{
				Prompt:            "\033[31m»\033[0m ",
				HistoryFile:       "/tmp/readline.tmp",
				InterruptPrompt:   "^C",
				EOFPrompt:         "exit",
				HistorySearchFold: true,
			})
			if err != nil {
				fmt.Printf("error setting up readline: %v\n", err)
				return
			}
			defer l.Close()
			l.CaptureExitSignal()

			for i := 0; i < 6; i++ {

				line, err := l.Readline()
				if err != nil {
					fmt.Printf("error reading line: %v\n", err)
					return
				}

				split := strings.Split(line, " ")
				if len(split) != 2 {
					fmt.Printf("invalid input: %s\n", line)
					fmt.Printf("input should be in the format '{word} {info} e.g. stray 22100'\n")
					return
				}
				word := split[0]
				info := split[1]
				err = criteria.Update(word, info)
				if err != nil {
					fmt.Printf("error updating criteria: %v\n", err)
				}

				entropy = wordle.Search(entropy, criteria)
				words := maps.Keys(entropy)
				// Sort the slice based on the value
				sort.Slice(words, func(i, j int) bool {
					return entropy[words[i]] < entropy[words[j]]
				})

				for _, word := range words {
					entropyVal := strconv.FormatFloat(entropy[word], 'f', 2, 64)
					fmt.Printf("%s: %s\n", word, entropyVal)
				}
			}

		},
	}

	return cmd
}

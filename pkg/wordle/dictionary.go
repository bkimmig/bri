package wordle

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var (
	macOSDict   = "/usr/share/dict/words"
	letterCount = 5

	MacOSType     = "darwin" // macos
	WindowsOSType = "windows"
)

func loadWords() ([]string, error) {
	osType := runtime.GOOS

	switch osType {
	case MacOSType:
		return loadDictionary(macOSDict)
	default:
		return nil, fmt.Errorf("unsupported os type: %s", osType)
	}
}

func loadDictionary(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == letterCount {
			lines = append(lines, strings.ToLower(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

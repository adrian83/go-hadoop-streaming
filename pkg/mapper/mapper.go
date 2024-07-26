package mapper

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var runesToRemove = []string{".", ",", ":", ";", "\"", "'", "?", "!", "(", ")", "-", "\\", "_"}

// Map maps data from source (io.Reader) and writes result to destination (io.Writer).
func Map(source io.Reader, destination io.Writer) {
	input := bufio.NewScanner(source)

	for input.Scan() {
		line := input.Text()
		words := strings.Split(line, " ")
		for _, word := range words {
			if word := normalize(word); word != "" {
				outputLine := fmt.Sprintf("%s\t1\n", word)
				fmt.Fprint(destination, outputLine)
			}
		}
	}
}

func normalize(word string) string {
	if word == "" {
		return word
	}

	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	for _, ch := range runesToRemove {
		word = strings.Replace(word, ch, "", -1)
	}

	return word
}

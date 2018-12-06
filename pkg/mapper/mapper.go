package mapper

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var toRemove = []string{".", ",", ":", ";", "\"", "'", "?", "!", "(", ")", "-", "\\", "_"}

// Map maps data from source (io.Reader) and writes result to destination (io.Writer).
func Map(source io.Reader, destination io.Writer) {
	input := bufio.NewScanner(source)

	for input.Scan() {
		line := input.Text()
		words := strings.Split(line, " ")
		for _, word := range words {
			if word := normalize(word); word != "" {
				fmt.Fprintln(destination, fmt.Sprintf("%s\t1\n", word))
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
	word = removeUselessRunes(word)
	return word
}

func removeUselessRunes(word string) string {
	correct := true
	for _, ch := range toRemove {
		if strings.HasPrefix(word, ch) {
			word = strings.TrimPrefix(word, ch)
			correct = false
		}
		if strings.HasSuffix(word, ch) {
			word = strings.TrimSuffix(word, ch)
			correct = false
		}
	}
	if correct {
		return word
	}
	return removeUselessRunes(word)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	toRemove = []string{".", ",", ":", ";", "\"", "'", "?", "!", "(", ")"}
)

func main() {

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		words := strings.Split(line, " ")
		for _, word := range words {
			if word := normalize(word); word != "" {
				fmt.Printf("%s\t1\n", word)
			}
		}
	}
}

func normalize(word string) string {
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

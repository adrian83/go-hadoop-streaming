package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	invalid = []string{".", ",", ":", ";", "\"", "'", "?", "!", "(", ")"}
)

func main() {

	bio := bufio.NewReader(os.Stdin)

	for true {

		line, err := bio.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while reading from Stdin. Error: %v\n", err)
			break
		}

		words := strings.Split(line, " ")
		for _, word := range words {
			if word, ok := normalize(word); ok {
				fmt.Printf("%s\t1\n", word)
			}
		}

	}
}

func normalize(word string) (string, bool) {
	word = strings.TrimSpace(word)
	if word == "" {
		return "", false
	}
	word = strings.ToLower(word)
	word = removeInvalidChar(word)
	if word == "" {
		return "", false
	}
	return word, true
}

func removeInvalidChar(word string) string {
	hasInvalidChar := false

	for _, ch := range invalid {
		hasInvalidChar = hasInvalidChar || strings.HasPrefix(word, ch) || strings.HasSuffix(word, ch)
		word = strings.TrimPrefix(word, ch)
		word = strings.TrimSuffix(word, ch)
	}

	if hasInvalidChar {
		word = removeInvalidChar(word)
	}

	return word
}

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

		if line = strings.TrimSpace(line); line != "" {
			words := strings.Split(line, " ")

			for _, word := range words {
				if word = normalize(word); word != "" {
					fmt.Println(fmt.Sprintf("%s\t1", word))
				}
			}
		}
	}
}

func normalize(word string) string {
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)
	return removeInvalidChar(word)
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

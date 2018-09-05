package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	bio := bufio.NewReader(os.Stdin)
	counts := make(map[string]int)

	for true {

		line, err := bio.ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while reading from Stdin. Error: %v\n", err)
			break
		}

		if line = strings.TrimSpace(line); line != "" {
			parts := strings.Split(line, "\t")

			word := parts[0]
			countStr := parts[1]

			count, err := strconv.Atoi(countStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error while processing data. Cannot convert string to number. Error: %v/n", err)
				break
			}
			counts[word] = counts[word] + count
		}
	}

	for word, count := range counts {
		fmt.Printf("%s\t%d\n", word, count)
	}
}

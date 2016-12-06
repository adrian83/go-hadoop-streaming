package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	bio := bufio.NewReader(os.Stdin)

	for true {

		line, err := bio.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error while reading from Stdin. Error: %v\n", err)
			}
			break
		}

		if line = strings.TrimSpace(line); line != "" {
			words := strings.Split(line, " ")

			for _, word := range words {
				word = strings.TrimSpace(word)
				if word != "" {
					fmt.Println(fmt.Sprintf("%s\t1", word))
				}
			}
		}
	}

}

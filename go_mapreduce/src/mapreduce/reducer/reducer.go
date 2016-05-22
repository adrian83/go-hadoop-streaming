package main

import (
	"bufio"
	"fmt"
	"io"
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
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "error: can't read - %s\n", err)
			}
			// io.EOF
			break

		}

		line = strings.TrimSpace(line)
		if len(line) > 0 {
			parts := strings.Split(line, "\t")

			word := parts[0]
			countStr := parts[1]

			count, err := strconv.Atoi(countStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "not a number! ", err)

				break
			}
			counts[word] = counts[word] + count

		}
	}
	for word, count := range counts {
		fmt.Printf("%s\t%d\n", word, count)
	}
	//fmt.Printf("%s\t%d\n", "test", 8)
}

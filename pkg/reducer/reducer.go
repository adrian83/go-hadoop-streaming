package reducer

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Reduce reduces data from source (io.Reader) and writes it to destination (io.Writer) or if error occurs to errors (io.Writer).
func Reduce(source io.Reader, errors, destination io.Writer) {
	input := bufio.NewScanner(source)
	counts := make(map[string]int)

	for input.Scan() {

		line := input.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "\t")

		word := parts[0]
		countStr := parts[1]

		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Fprintf(errors, "Error while processing data. Cannot convert string '%s' to number. Error: %v/n", countStr, err)
			break
		}
		counts[word] += count

	}

	for word, count := range counts {
		fmt.Fprintf(destination, fmt.Sprintf("%s\t%d\n", word, count))
	}
}

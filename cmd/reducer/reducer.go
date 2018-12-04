package main

import (
	"os"

	"github.com/adrian83/go-hadoop-streaming/pkg/reducer"
)

func main() {
	source := os.Stdin
	destination := os.Stdout
	errors := os.Stderr

	reducer.Reduce(source, errors, destination)
}

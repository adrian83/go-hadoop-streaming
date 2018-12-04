package main

import (
	"os"

	"github.com/adrian83/go-hadoop-streaming/pkg/mapper"
)

func main() {
	source := os.Stdin
	destination := os.Stdout

	mapper.Map(source, destination)
}

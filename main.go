package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andyklimenko/html-element-finder/src/parser"
)

func main() {
	idToFind := flag.String("id", "", "")
	originalPagePath := flag.String("original", "", "")
	updatedPagePath := flag.String("updated", "", "")
	flag.Parse()

	originalFile, err := os.Open(*originalPagePath)
	if err != nil {
		panic(err)
	}

	updatedFile, err := os.Open(*updatedPagePath)
	if err != nil {
		panic(err)
	}

	originalElementPath, node, err := parser.FindOriginalElement(originalFile, *idToFind)
	if err != nil {
		panic(err)
	}
	fmt.Println("path to element in original page: ", originalElementPath)

	foundElementPath, err := parser.FindMatch(updatedFile, node)
	if err != nil {
		panic(err)
	}

	fmt.Println("path to element in updated page: ", foundElementPath)
}

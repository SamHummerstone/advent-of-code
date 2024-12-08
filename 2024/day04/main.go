package main

import (
	"embed"
	"fmt"
)

//go:embed inputFile.txt
var inputFile embed.FS

func main() {
	inputFile, err := inputFile.ReadFile("inputFile.txt")
	if err != nil {
		panic("Could not read input file")
	}
	fmt.Println(Part1(inputFile))
	fmt.Println(Part2(inputFile))
}

func Part1(input []byte) int {
	return 0
}

func Part2(input []byte) int {
	return 0
}

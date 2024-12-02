package main

import (
	"embed"
	"fmt"
)

//go:embed input.txt
var inputFile embed.FS

func main() {
	inputString, err := inputFile.ReadFile("input.txt")
	if err != nil {
		panic("Could not load input file")
	}
	fmt.Println(Part1(string(inputString)))
}

func Part1(reports string) int {
	return 0
}

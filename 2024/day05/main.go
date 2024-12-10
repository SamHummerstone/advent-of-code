package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed inputFile.txt
var inputFile embed.FS

func main() {
	inputFile, err := inputFile.ReadFile("inputFile.txt")
	if err != nil {
		panic("Could not read input file")
	}

	inputSplit := strings.Split(string(inputFile), "\n\n")

	rules, pages := inputSplit[0], inputSplit[1]

	fmt.Println(Part1(rules, pages))
}

func Part1(rules, pages string) int {
	fmt.Println(rules)

	return 0
}

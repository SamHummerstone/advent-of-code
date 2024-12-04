package main

import (
	"embed"
	"fmt"
	"regexp"
)

//go:embed inputFile.txt
var inputFile embed.FS

func main() {
	inputFile, err := inputFile.ReadFile("inputFile.txt")
	if err != nil {
		panic("Could not read input file")
	}
	fmt.Println(Part1(inputFile))
}

func Part1(input []byte) int {
	var count int
	muls := ExtractOccurances(`(mul\([0-9]+,[0-9]+\))*`, string(input))

	for _, m := range muls {
		count += m.Equals()
	}

	return count
}

func ExtractOccurances(regex, source string) []Mul {
	var occurances []Mul
	re := regexp.MustCompile(regex)

	matches := re.FindStringSubmatch(source)

	return occurances
}

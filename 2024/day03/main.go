package main

import (
	"embed"
	"fmt"
	"regexp"
)

//go:embed inputFile.txt
var inputFile embed.FS
var mulExpr = `mul\(([0-9]+),([0-9]+)\)`

func main() {
	inputFile, err := inputFile.ReadFile("inputFile.txt")
	if err != nil {
		panic("Could not read input file")
	}
	fmt.Println(Part1(inputFile))
}

func Part1(input []byte) int {
	var count int
	var muls []Mul

	mulStrings := ExtractOccurances(mulExpr, string(input))

	fmt.Println(mulStrings)

	for _, m := range mulStrings {
		nm, err := New(m)
		if err != nil {
			fmt.Println("Could not convert string:", m)
			break
		}
		muls = append(muls, nm)
	}

	for _, m := range muls {
		count += m.Equals()
	}

	return count
}

func ExtractOccurances(regex, source string) []string {
	re := regexp.MustCompile(regex)
	matches := re.FindAllString(source, -1)

	return matches
}

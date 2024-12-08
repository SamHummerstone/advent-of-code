package main

import (
	"embed"
	"fmt"
	"regexp"
)

//go:embed inputFile.txt
var inputFile embed.FS
var mulExpr = `mul\(([0-9]+),([0-9]+)\)`
var mulDoDontExpr = `(mul\([0-9]+,[0-9]+\)|(do\(\))|(don't\(\)))?`

func main() {
	inputFile, err := inputFile.ReadFile("inputFile.txt")
	if err != nil {
		panic("Could not read input file")
	}
	fmt.Println(Part1(inputFile))
	fmt.Println(Part2(inputFile))
}

func Part1(input []byte) int {

	mulStrings := ExtractOccurances(mulExpr, string(input))

	return GetMulsCount(mulStrings)

}

func Part2(input []byte) int {
	mulStrings := ExtractDoOccurances(mulDoDontExpr, string(input))

	fmt.Println(mulStrings)

	return GetMulsCount(mulStrings)

}

func GetMulsCount(mulStrings []string) int {
	var muls []Mul
	var count int
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

func ExtractDoOccurances(regex, source string) []string {
	var matches []string
	enabled := true
	re := regexp.MustCompile(regex)
	grossMatches := re.FindAllString(source, -1)

	for _, match := range grossMatches {
		switch match {
		case "do()":
			{
				enabled = true
			}
		case "don't()":
			{
				enabled = false
			}
		default:
			{
				if enabled && match != "" {
					fmt.Println("Found 'Do' before", match)
					matches = append(matches, match)
				}
			}
		}
	}

	return matches
}

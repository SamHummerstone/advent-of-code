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

	ruleList, pageList := ProcessInputString(inputFile)

	fmt.Println(Part1(ruleList, pageList))
}

func Part1(rules []Rule, pages []Page) int {
	var result int

	for _, page := range pages {
		relevantRules := GetRelevantRules(rules, page)
		if page.IsValid(relevantRules) {
			// fmt.Println("page", page, "has relevant rules:", relevantRules)
			result += page.GetMiddleNum()
		}
	}

	return result
}

func ProcessInputString(input []byte) ([]Rule, []Page) {
	var ruleList []Rule
	var pageList []Page

	inputSplit := strings.Split(string(input), "\n\n")

	rules, pages := inputSplit[0], inputSplit[1]

	rulesStrings := strings.Split(rules, "\n")
	for _, rule := range rulesStrings {
		ruleList = append(ruleList, NewRule(rule))
	}

	pageStrings := strings.Split(pages, "\n")
	for _, p := range pageStrings {
		pageList = append(pageList, NewPage(p))
	}

	return ruleList, pageList
}

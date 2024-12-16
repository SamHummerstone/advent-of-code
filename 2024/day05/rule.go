package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	Left  int
	Right int
}

func NewRule(s string) Rule {
	split := strings.Split(s, "|")
	l, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("Could not convert string to int:", split[0])
		panic(err)
	}
	r, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println("Could not convert string to int:", split[1])
		panic(err)
	}

	return Rule{
		Left:  l,
		Right: r,
	}
}

func GetRelevantRules(rules []Rule, page Page) []Rule {
	var relevantRules []Rule

	for _, rule := range rules {
		l, _ := page.Contains(rule.Left)
		r, _ := page.Contains(rule.Right)
		if l && r {
			relevantRules = append(relevantRules, rule)
		}
	}

	return relevantRules
}

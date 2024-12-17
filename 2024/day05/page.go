package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Page []int

func NewPage(s string) Page {
	var page Page
	split := strings.Split(s, ",")
	for _, numstring := range split {
		i, err := strconv.Atoi(numstring)
		if err != nil {
			fmt.Println("Could not convert to integer:", numstring)
			return nil
		}
		page = append(page, i)
	}

	return page
}

func (p Page) GetMiddleNum() int {
	size := len(p)
	if size == 0 {
		return 0
	}
	mid := ((size + 1) / 2) - 1

	return p[mid]
}

func (p Page) CheckRule(rule Rule) bool {
	_, l := p.Contains(rule.Left)
	_, r := p.Contains(rule.Right)

	return l < r
}

func (p Page) Contains(n int) (bool, int) {
	for i, pageNum := range p {
		if pageNum == n {
			return true, i
		}
	}
	return false, -1
}

func (p Page) IsValid(rules []Rule) bool {
	// Rules are all relavant
	for _, rule := range rules {
		if !p.CheckRule(rule) {
			return false
		}
	}
	return true
}

func (p Page) Reorder(relevantRules []Rule) Page {
	var newPage Page = p

	return newPage
}

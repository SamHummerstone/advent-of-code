package main

import (
	"testing"
)

var testString = []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

var ruleTestList []Rule
var pageTestList []Page

func init() {
	ruleTestList, pageTestList = ProcessInputString(testString)
}

func Test_Part1(t *testing.T) {
	t.Run("Test Part1", func(t *testing.T) {
		want := 143
		got := Part1(ruleTestList, pageTestList)

		if want != got {
			t.Errorf("Wanted: %v\nGot: %v", want, got)
		}
	})
}

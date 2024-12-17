package main

import (
	"slices"
	"testing"
)

func Test_GetRelevantRules(t *testing.T) {
	page := Page{
		75,
		47,
		61,
		53,
		29,
	}
	want := []Rule{
		{Left: 47, Right: 53},
		{Left: 75, Right: 29},
		{Left: 75, Right: 53},
		{Left: 53, Right: 29},
		{Left: 61, Right: 53},
		{Left: 61, Right: 29},
		{Left: 75, Right: 47},
		{Left: 47, Right: 61},
		{Left: 75, Right: 61},
		{Left: 47, Right: 29},
	}

	t.Run("Test GetRelevant Rules", func(t *testing.T) {
		got := GetRelevantRules(ruleTestList, page)

		if !slices.Equal(got, want) {
			t.Errorf("Wanted: %v\nGot: %v", want, got)
		}
	})
}

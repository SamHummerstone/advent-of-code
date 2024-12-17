package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestPage_GetMiddleNum_EvenSize(t *testing.T) {
	page := Page{1, 2, 3}
	mid := page.GetMiddleNum()
	if mid != 2 {
		t.Errorf("Expected middle num to be 2, but got %d", mid)
	}
}

func TestPage_GetMiddleNum_EmptyPage(t *testing.T) {
	page := Page{}
	mid := page.GetMiddleNum()
	if mid != 0 {
		t.Errorf("Expected middle num to be 0 for empty page, but got %d", mid)
	}
}

func TestPage_GetMiddleNum_SingleElementPage(t *testing.T) {
	page := Page{5}
	mid := page.GetMiddleNum()
	if mid != 5 {
		t.Errorf("Expected middle num to be 5 for single-element page, but got %d", mid)
	}
}

func TestPage_IsValid(t *testing.T) {
	tcs := []struct {
		Input  Page
		Expect bool
	}{
		{
			Input:  pageTestList[0],
			Expect: true,
		},
		{
			Input:  pageTestList[1],
			Expect: true,
		},
		{
			Input:  pageTestList[2],
			Expect: true,
		},
		{
			Input:  pageTestList[3],
			Expect: false,
		},
		{
			Input:  pageTestList[4],
			Expect: false,
		},
		{
			Input:  pageTestList[5],
			Expect: false,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test IsValid - Case %v", i+1), func(t *testing.T) {
			relevantRules := GetRelevantRules(ruleTestList, tc.Input)
			got := tc.Input.IsValid(relevantRules)
			want := tc.Expect
			if got != want {
				t.Errorf("Got: %v\nWanted: %v", got, want)
			}
		})
	}
}

func TestPage_Reorder(t *testing.T) {
	tcs := []struct {
		Input  Page
		Expect Page
	}{
		{
			Input: pageTestList[3],
			Expect: Page{
				97, 75, 47, 61, 53,
			},
		},
		{
			Input: pageTestList[4],
			Expect: Page{
				61, 29, 13,
			},
		},
		{
			Input: pageTestList[5],
			Expect: Page{
				97, 75, 47, 29, 13,
			},
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test Reorder - Case %v", i+1), func(t *testing.T) {
			relevantRules := GetRelevantRules(ruleTestList, tc.Input)
			want := tc.Expect
			got := tc.Input.Reorder(relevantRules)

			if !slices.Equal(got, want) {
				t.Errorf("Wanted: %v\nGot: %v", want, got)
			}
		})
	}
}

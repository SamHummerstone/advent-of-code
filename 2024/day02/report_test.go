package main

import (
	"fmt"
	"slices"
	"testing"
)

func Test_New(t *testing.T) {
	tcs := []struct {
		Input  []int
		Expect Report
	}{
		{
			Input: []int{7, 6, 4, 2, 1},
			Expect: Report{
				Content: []int{7, 6, 4, 2, 1},
				Safe:    true,
			},
		},
		{
			Input: []int{1, 2, 7, 8, 9},
			Expect: Report{
				Content: []int{1, 2, 7, 8, 9},
				Safe:    false,
			},
		},
		{
			Input: []int{9, 7, 6, 2, 1},
			Expect: Report{
				Content: []int{9, 7, 6, 2, 1},
				Safe:    false,
			},
		},
		{
			Input: []int{1, 3, 2, 4, 5},
			Expect: Report{
				Content: []int{1, 3, 2, 4, 5},
				Safe:    false,
			},
		},
		{
			Input: []int{8, 6, 4, 4, 1},
			Expect: Report{
				Content: []int{8, 6, 4, 4, 1},
				Safe:    false,
			},
		},
		{
			Input: []int{1, 3, 6, 7, 9},
			Expect: Report{
				Content: []int{1, 3, 6, 7, 9},
				Safe:    true,
			},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test New - Case %v", i), func(t *testing.T) {
			want := tc.Expect
			got := New(tc.Input)
			if slices.Equal(want.Content, got.Content) && want.Safe != got.Safe {
				t.Errorf("Wanted: %v\nGot: %v", want, got)
			}
		})
	}
}

func Test_IsSafe(t *testing.T) {
	tcs := []struct {
		Input  []int
		Expect bool
	}{
		{
			Input:  []int{7, 6, 4, 2, 1},
			Expect: true,
		},
		{
			Input:  []int{1, 2, 7, 8, 9},
			Expect: false,
		},
		{
			Input:  []int{9, 7, 6, 2, 1},
			Expect: false,
		},
		{
			Input:  []int{1, 3, 2, 4, 5},
			Expect: false,
		},
		{
			Input:  []int{8, 6, 4, 4, 1},
			Expect: false,
		},
		{
			Input:  []int{1, 3, 6, 7, 9},
			Expect: true,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test IsReportSafe - Case %v", i), func(t *testing.T) {
			want := tc.Expect
			got := Report{Content: tc.Input}
			if want != got.IsSafe() {
				t.Errorf("Wanted: %v\nGot: %v", want, got)
			}
		})
	}
}

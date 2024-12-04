package main

import (
	"fmt"
	"slices"
	"testing"
)

var (
	TestString = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
)

func Test_Part1(t *testing.T) {
	t.Run("Test Part1 func", func(t *testing.T) {
		want := 2
		got := Part1(TestString)

		if want != got {
			t.Errorf("Wanted: %v\nGot: %v\n", want, got)
		}
	})
}

func Test_Part2(t *testing.T) {
	t.Run("Test Part1 func", func(t *testing.T) {
		want := 4
		got := Part2(TestString)

		if want != got {
			t.Errorf("Wanted: %v\nGot: %v\n", want, got)
		}
	})
}

// want1 := []int{7, 4, 2, 1}
// want2 := []int{7, 6, 2, 1}
// want3 := []int{7, 6, 4, 1}
// want4 := []int{7, 6, 4, 2}

func Test_GetDampenedReports(t *testing.T) {
	// t.Run("Test GetDampenedReports func - Case 0", func(t *testing.T) {
	// 	tc := []int{7, 6, 4, 2, 1}
	// 	want := []int{6, 4, 2, 1}

	// 	got := GetDampenedReport(tc, 0)

	// 	if !slices.Equal(want, got) {
	// 		t.Errorf("Wanted: %v\nGot: %v", want, got)
	// 	}
	// })

}

func Test_ReportStringToInts(t *testing.T) {
	tcs := []struct {
		Input  string
		Expect []int
	}{
		{
			Input:  "7 6 4 2 1",
			Expect: []int{7, 6, 4, 2, 1},
		},
		{
			Input:  "1 2 7 8 9",
			Expect: []int{1, 2, 7, 8, 9},
		},
		{
			Input:  "9 7 6 2 1",
			Expect: []int{9, 7, 6, 2, 1},
		},
		{
			Input:  "1 3 2 4 5",
			Expect: []int{1, 3, 2, 4, 5},
		},
		{
			Input:  "8 6 4 4 1",
			Expect: []int{8, 6, 4, 4, 1},
		},
		{
			Input:  "1 3 6 7 9",
			Expect: []int{1, 3, 6, 7, 9},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Testing ReportStringToInts - Case %v", i), func(t *testing.T) {
			want := tc.Expect
			got := ReportStringToInts(tc.Input)
			if !slices.Equal(want, got) {
				t.Errorf("Wanted: %v\nGot: %v", want, got)
			}
		})
	}
}

package main

import (
	"fmt"
	"testing"
)

var testInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_Part1(t *testing.T) {
	want := 18
	t.Run("Test Part1", func(t *testing.T) {
		got := Part1(testInput)
		if got != want {
			t.Errorf("Got %v:\nWant: %v", got, want)
		}
	})
}

func Test_SearchX(t *testing.T) {
	match := []string{"X", "M", "A", "S"}
	tcs := []struct {
		Input     []string
		Direction int
		Expect    bool
	}{
		{
			Input:     match,
			Direction: 1,
			Expect:    true,
		},
		{
			Input:     match,
			Direction: -1,
			Expect:    false,
		},
		{
			Input:     []string{"S", "A", "M", "X"},
			Direction: 1,
			Expect:    false,
		},
		{
			Input:     []string{"S", "A", "M", "X"},
			Direction: -1,
			Expect:    true,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test SearchX - case %v", i), func(t *testing.T) {
			want := tc.Expect
			var got bool
			if tc.Direction < 0 {
				got = SearchX([][]string{tc.Input}, match, len(tc.Input)-1, 0, tc.Direction)
			} else {
				got = SearchX([][]string{tc.Input}, match, 0, 0, tc.Direction)
			}
			if want != got {
				t.Errorf("Want: %v\nGot: %v", want, got)
			}
		})
	}
}

func Test_SearchY(t *testing.T) {
	match := []string{"X", "M", "A", "S"}
	tcs := []struct {
		Input     []string
		Direction int
		Expect    bool
	}{
		{
			Input:     match,
			Direction: 1,
			Expect:    true,
		},
		{
			Input:     match,
			Direction: -1,
			Expect:    false,
		},
		{
			Input:     []string{"S", "A", "M", "X"},
			Direction: 1,
			Expect:    false,
		},
		{
			Input:     []string{"S", "A", "M", "X"},
			Direction: -1,
			Expect:    true,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test SearchY - case %v", i), func(t *testing.T) {
			want := tc.Expect
			var got bool
			if tc.Direction < 0 {
				got = SearchY([][]string{tc.Input}, match, 0, len(tc.Input)-1, tc.Direction)
			} else {
				got = SearchY([][]string{tc.Input}, match, 0, 0, tc.Direction)
			}
			if want != got {
				t.Errorf("Want: %v\nGot: %v", want, got)
			}
		})
	}
}

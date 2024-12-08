package main

import (
	"slices"
	"testing"
)

var (
	testString  = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	testString2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
)

func Test_Part1(t *testing.T) {
	t.Run("Test Part1", func(t *testing.T) {
		want := 161
		got := Part1([]byte(testString))
		if want != got {
			t.Errorf("Wanted: %v\nGot: %v", want, got)
		}
	})
}

func Test_Part2(t *testing.T) {
	t.Run("Test Part2", func(t *testing.T) {
		want := 48
		got := Part2([]byte(testString2))
		if want != got {
			t.Errorf("Wanted: %v\nGot: %v", want, got)
		}
	})
}

func Test_ExtractOccurances(t *testing.T) {
	t.Run("Test ExtractOccurances", func(t *testing.T) {
		want := []string{
			"mul(2,4)",
			"mul(5,5)",
			"mul(11,8)",
			"mul(8,5)",
		}
		got := ExtractOccurances(mulExpr, testString)

		if !slices.Equal(got, want) {
			t.Errorf("Wanted: %v\nGot: %v", want, got)
		}
	})
}

func Test_ExtractDoOccurances(t *testing.T) {
	t.Run("Test ExtractDoOccurances", func(t *testing.T) {
		want := []string{
			"mul(2,4)",
			"mul(8,5)",
		}
		got := ExtractDoOccurances(mulDoDontExpr, testString2)

		if !slices.Equal(got, want) {
			t.Errorf("Wanted: %v\nGot: %v", want, got)
		}
	})
}

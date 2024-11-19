package main

import (
	"strings"
	"testing"
)

var example1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var example2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_ConvertWrittenNumber(t *testing.T) {
	tcs := []struct {
		input  string
		expect string
	}{
		{
			input:  "two1nine",
			expect: "t2o1n9e",
		},
		{
			input:  "eightwothree",
			expect: "e8t2ot3e",
		},
	}

	for _, tc := range tcs {
		t.Run("Test Converting written numbers", func(t *testing.T) {
			if res := ConvertWrittenNumber(tc.input); res != tc.expect {
				t.Errorf("Want: %v\nGot: %v", tc.expect, res)
			}
		})
	}
}

func Test_Part1(t *testing.T) {
	lines := strings.Split(string(example1), "\n")
	t.Run("Day 1 - Part 1", func(t *testing.T) {
		want := 142
		got := Part1(lines)

		if got != want {
			t.Errorf("Want: %v\nGot: %v", want, got)
		}
	})
}

func Test_Part2(t *testing.T) {
	lines := strings.Split(string(example2), "\n")
	t.Run("Day 1 - Part 2", func(t *testing.T) {
		want := 281
		got := Part2(lines)

		if got != want {
			t.Errorf("Want: %v\nGot: %v", want, got)
		}
	})
}

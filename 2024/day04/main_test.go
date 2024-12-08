package main

import "testing"

var testInput = []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

func Test_Part1(t *testing.T) {
	want := 18
	t.Run("Test Part1", func(t *testing.T) {
		got := Part1(testInput)
		if got != want {
			t.Errorf("Got %v:\nWant: %v", got, want)
		}
	})
}

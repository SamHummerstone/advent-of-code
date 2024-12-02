package main

import "testing"

var (
	TestString = `
7 6 4 2 1
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

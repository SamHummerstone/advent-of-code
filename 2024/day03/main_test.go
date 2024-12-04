package main

import "testing"

var (
	testString = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
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

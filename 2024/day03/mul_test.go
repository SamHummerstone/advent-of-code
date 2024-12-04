package main

import (
	"fmt"
	"testing"
)

func Test_Mul_Equals(t *testing.T) {
	tcs := []struct {
		Input  Mul
		Expect int
	}{
		{
			Input: Mul{
				Left:  2,
				Right: 4,
			},
			Expect: 8,
		},
		{
			Input: Mul{
				Left:  5,
				Right: 5,
			},
			Expect: 25,
		},
		{
			Input: Mul{
				Left:  11,
				Right: 8,
			},
			Expect: 88,
		},
		{
			Input: Mul{
				Left:  8,
				Right: 5,
			},
			Expect: 40,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test Mul.Equals() - Case %v", i), func(t *testing.T) {
			want := tc.Expect
			got := tc.Input.Equals()

			if want != got {
				t.Errorf("Wanted: %v\nGot: %v", want, got)
			}
		})
	}
}

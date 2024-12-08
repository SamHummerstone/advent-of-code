package main

import (
	"fmt"
	"testing"
)

func Test_Equals(t *testing.T) {
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

// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
func Test_New(t *testing.T) {
	tcs := []struct {
		Input  string
		Expect Mul
	}{
		{
			Input: "mul(2,4)",
			Expect: Mul{
				Left:  2,
				Right: 4,
			},
		},
		{
			Input: "mul(5,5)",
			Expect: Mul{
				Left:  5,
				Right: 5,
			},
		},
		{
			Input: "mul(11,8)",
			Expect: Mul{
				Left:  11,
				Right: 8,
			},
		},
		{
			Input: "mul(8,5)",
			Expect: Mul{
				Left:  8,
				Right: 5,
			},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Testing New - Case %v", i), func(t *testing.T) {
			want := tc.Expect
			got, err := New(tc.Input)
			if err != nil || want != got {
				t.Errorf("Could not create Mul from string\nwant: %v\ngot: %v", want, got)
			}
		})
	}
}

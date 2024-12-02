package main

import (
	"fmt"
	"testing"
)

func Test_GetDifference(t *testing.T) {
	tcs := []struct {
		input  [2]int
		expect int
	}{
		{
			input:  [2]int{3, 4},
			expect: 1,
		},
		{
			input:  [2]int{4, 3},
			expect: 1,
		},
		{
			input:  [2]int{2, 5},
			expect: 3,
		},
		{
			input:  [2]int{1, 3},
			expect: 2,
		},
		{
			input:  [2]int{3, 9},
			expect: 6,
		},
		{
			input:  [2]int{3, 3},
			expect: 0,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test GetDifference case #%v", i), func(t *testing.T) {
			got := GetDifference(tc.input[0], tc.input[1])
			if tc.expect != got {
				t.Errorf("Want: %v \nGot: %v", tc.expect, got)
			}
		})
	}
}

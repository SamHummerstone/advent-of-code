package main

import "testing"

func Test_New(t *testing.T) {
	t.Run("Interpret Game string", func(t *testing.T) {
		tcs := []struct {
			input  string
			expect Game
		}{
			{
				input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				expect: Game{
					Id:    1,
					Blue:  6,
					Red:   4,
					Green: 2,
				},
			},
			{
				input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				expect: Game{
					Id:    2,
					Blue:  4,
					Red:   1,
					Green: 3,
				},
			},
			{
				input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				expect: Game{
					Id:    3,
					Red:   20,
					Green: 13,
					Blue:  6,
				},
			},
			{
				input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				expect: Game{
					Id:    4,
					Red:   14,
					Green: 3,
					Blue:  15,
				},
			},
		}
		for _, tc := range tcs {
			if res := New(tc.input); *res != tc.expect {
				t.Errorf("Want: %v\nGot: %v", tc.expect, *res)
			}
		}
	})
}

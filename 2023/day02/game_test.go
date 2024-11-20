package main

import (
	"strings"
	"testing"
)

func Test_New(t *testing.T) {
	t.Run("Interpret Game string", func(t *testing.T) {
		gs := LinestoGames(strings.Split(examples1, "\n"))
		tcs := []struct {
			input  string
			expect Game
		}{
			{
				input:  "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				expect: *gs[0],
			},
			{
				input:  "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				expect: *gs[1],
			},
			{
				input:  "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				expect: *gs[2],
			},
			{
				input:  "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				expect: *gs[3],
			},
			{
				input:  "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
				expect: *gs[4],
			},
		}
		for _, tc := range tcs {
			if res := New(tc.input); *res != tc.expect {
				t.Errorf("Want: %v\nGot: %v", tc.expect, *res)
			}
		}
	})
}

func Test_CheckGame(t *testing.T) {
	t.Run("Test Game Checking", func(t *testing.T) {
		gs := LinestoGames(strings.Split(examples1, "\n"))

		tcs := []struct {
			input  Game
			expect bool
		}{
			{
				input:  *gs[0],
				expect: true,
			},
			{
				input:  *gs[1],
				expect: true,
			},
			{
				input:  *gs[2],
				expect: false,
			},
			{
				input:  *gs[3],
				expect: false,
			},
			{
				input:  *gs[4],
				expect: true,
			},
		}
		for _, tc := range tcs {
			if res := tc.input.CheckGame(gc); res != tc.expect {
				t.Errorf("Want: %v\nGot: %v", tc.expect, res)
			}
		}
	})
}

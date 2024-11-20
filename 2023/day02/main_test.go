package main

import (
	"strings"
	"testing"
)

var examples1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var gc = GameDice{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func Test_Part1(t *testing.T) {
	t.Run("Test Part1", func(t *testing.T) {
		games := LinestoGames(strings.Split(examples1, "\n"))

		want := 8
		got := Part1(games, gc)

		if got != want {
			t.Errorf("Want: %v\nGot: %v", want, got)
		}
	})
}

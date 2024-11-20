package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Red   int
	Green int
	Blue  int
}

type GameDice struct {
	Red   int
	Green int
	Blue  int
}

func New(s string) *Game {
	var red, green, blue int = 0, 0, 0
	idStr, rem, _ := strings.Cut(s, ":")
	id, err := strconv.Atoi(strings.TrimPrefix(idStr, "Game "))
	if err != nil {
		fmt.Println("Could not convert ID to int:", err)
		os.Exit(1)
	}

	// Split into rounds
	for _, r := range strings.Split(rem, ";") {
		// Split into colours
		for _, c := range strings.Split(r, ",") {
			colourCount := strings.Split(strings.Trim(c, " "), " ")
			count, err := strconv.Atoi(colourCount[0])
			if err != nil {
				fmt.Printf("Cannot convert str to int: %v", err)
				os.Exit(1)
			}
			switch colour := colourCount[1]; colour {
			case "red":
				if count > red {
					red = count
				}
			case "blue":
				if count > blue {
					blue = count
				}
			case "green":
				if count > green {
					green = count
				}
			}
		}
	}
	return &Game{
		Id:    id,
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func (g Game) CheckGame(check GameDice) bool {
	switch {
	case g.Red > check.Red:
		return false
	case g.Blue > check.Blue:
		return false
	case g.Green > check.Green:
		return false
	default:
		return true
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Colour struct {
	Max int
	// Min int // Not needed, I was overcomplicating things
}

type Game struct {
	Id    int
	Red   Colour
	Green Colour
	Blue  Colour
}

type GameLimits struct {
	Red   int
	Green int
	Blue  int
}

func New(s string) *Game {
	var red, green, blue Colour
	id, rem := IDFromString(s)

	// Split into rounds
	for _, r := range strings.Split(rem, ";") {
		// Split into colours
		for _, c := range strings.Split(r, ",") {
			count, colour := GetColourAndCount(c)
			switch colour {
			case "red":
				if count > red.Max {
					red.Max = count
				}
			case "blue":
				if count > blue.Max {
					blue.Max = count
				}
			case "green":
				if count > green.Max {
					green.Max = count
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

func (g Game) CheckGame(check GameLimits) bool {
	switch {
	case g.Red.Max > check.Red:
		return false
	case g.Blue.Max > check.Blue:
		return false
	case g.Green.Max > check.Green:
		return false
	default:
		return true
	}
}

func (g Game) Power() int {
	return g.Red.Max * g.Green.Max * g.Blue.Max
}

// Get game ID from the beginning of the string and
// return the ID and the remainder of the string
func IDFromString(s string) (int, string) {
	before, after, found := strings.Cut(s, ": ")
	if !found {
		fmt.Println("String doesn't include ': ' character")
		return 0, ""
	}
	id, err := strconv.Atoi(strings.TrimPrefix(before, "Game "))
	if err != nil {
		fmt.Println("IDFromString:", err)
	}

	return id, after
}

// Get colour and count from string
func GetColourAndCount(s string) (int, string) {
	colourCount := strings.Split(strings.Trim(s, " "), " ")
	count, err := strconv.Atoi(colourCount[0])
	if err != nil {
		fmt.Printf("Cannot convert str to int: %v", err)
		os.Exit(1)
	}
	colour := colourCount[1]

	return count, colour
}

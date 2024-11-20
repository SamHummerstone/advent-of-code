package main

import (
	"embed"
	"fmt"
	"log"
	"strings"
)

var (
	//go:embed input.txt
	inputFile embed.FS
)

func main() {
	fmt.Println("Cube Conundrum")

	input, err := inputFile.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Could not open input file:", err)
	}
	lines := strings.Split(string(input), "\n")

	games := LinestoGames(lines)
	gameDice := GameDice{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	solution1 := Part1(games, gameDice)

	fmt.Println(solution1)

}

func Part1(games []*Game, gd GameDice) int {
	var idSum int = 0
	for _, g := range games {
		if g.CheckGame(gd) {
			idSum += g.Id
		}
	}

	return idSum
}

func LinestoGames(lines []string) []*Game {
	// Make a list of games from the lines
	games := make([]*Game, len(lines))
	for i, l := range lines {
		games[i] = New(l)
	}
	return games
}

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
	gameLimits := GameLimits{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	solution1 := Part1(games, gameLimits)
	solution2 := Part2(games)

	fmt.Println(solution1)
	fmt.Println(solution2)

}

func Part1(games []*Game, gd GameLimits) int {
	var idSum int = 0
	for _, g := range games {
		if g.CheckGame(gd) {
			idSum += g.Id
		}
	}

	return idSum
}

func Part2(games []*Game) int {
	var powerSum int = 0
	for _, g := range games {
		powerSum += g.Power()
	}
	return powerSum
}

func LinestoGames(lines []string) []*Game {
	// Make a list of games from the lines
	games := make([]*Game, len(lines))
	for i, l := range lines {
		games[i] = New(l)
	}
	return games
}

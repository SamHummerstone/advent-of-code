package main

import (
	"embed"
	"fmt"
	"strings"
)

type Board [][]string

//go:embed inputFile.txt
var inputFile embed.FS

func main() {
	inputFile, err := inputFile.ReadFile("inputFile.txt")
	if err != nil {
		panic("Could not read input file")
	}
	fmt.Println(Part1(string(inputFile)))
	fmt.Println(Part2(inputFile))
}

func Part1(input string) int {
	match := []string{"X", "M", "A", "S"}
	board := CreateBoard(input)

	TraverseBoard(board, match, func() {
		fmt.Println("tada")
	})
	return 0
}

func Part2(input []byte) int {
	return 0
}

func CreateBoard(input string) Board {
	rows := strings.Split(string(input), "\n")
	board := make(Board, len(rows))

	for y, row := range rows {
		board[y] = make([]string, len(row))
		for x, r := range row {
			board[y][x] = string(r)
		}
	}
	return board
}

func TraverseBoard(board Board, match []string, f func()) {
	for y, row := range board {
		for x, col := range row {
			if col == match[0] {
				fmt.Printf("Found an %v at pos (%v,%v)", col, x, y)
				f()
			}
		}
	}
}

func SearchX(board [][]string, match []string, startX, startY int, directionX int) bool {
	for i := 0; i < len(match); i++ {
		if board[startY][startX+(i*directionX)] != match[i] {
			return false
		}
	}
	return true
}

func SearchY(board [][]string, match []string, startX, startY int, directionY int) bool {
	for i := 0; i < len(match); i++ {
		if board[startY+(i*directionY)][startX] != match[i] {
			return false
		}
	}
	return true
}

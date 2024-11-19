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

	for _, i := range lines {
		New(i)
	}

}

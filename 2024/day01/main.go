package main

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputFile embed.FS

func main() {
	fmt.Println(Part1())
	before := time.Now()
	fmt.Println(Part2(), "took", time.Since(before))
}

func Part1() int {
	var totalDistance int
	input, err := inputFile.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open file: %v", err)
	}

	left, right := SplitDistances(input)
	// Sanity check lengths
	if len(left) != len(right) {
		fmt.Printf("Left length doesn't match right length")
	}

	for i, n := range left {
		totalDistance += GetDifference(n, right[i])
	}

	return totalDistance
}

func Part2() int {
	var similarityScoreSum int
	input, err := inputFile.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open file: %v", err)
	}

	left, right := SplitDistances(input)
	// Sanity check lengths
	if len(left) != len(right) {
		fmt.Printf("Left length doesn't match right length")
	}

	for _, i := range left {
		similarityScoreSum += GetOccurancesMultiplied(i, right)
	}

	return similarityScoreSum
}

func SplitDistances(inputFile []byte) ([]int, []int) {
	var left, right []int

	lines := strings.Split(string(inputFile), "\n")
	for _, i := range lines {
		digits := strings.Split(i, "   ")
		l, err := strconv.Atoi(digits[0])
		if err != nil {
			fmt.Println("Cannot convert string to int:", err)
		}
		r, err := strconv.Atoi(digits[1])
		if err != nil {
			fmt.Println("Cannot convert string to int:", err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func GetDifference(left, right int) int {
	if difference := left - right; difference < 0 {
		return -difference
	} else {
		return difference
	}
}

func GetOccurancesMultiplied(n int, ns []int) int {
	var count int

	for _, i := range ns {
		if i == n {
			count++
		}
	}

	return count * n
}

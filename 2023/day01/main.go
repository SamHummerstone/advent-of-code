package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var (
	//go:embed input.txt
	inputFile embed.FS
)

func main() {
	input, err := inputFile.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Could not open input file:", err)
	}
	lines := strings.Split(string(input), "\n")

	log.Println("input list length:", len(lines))
	part1 := Part1(lines)
	part2 := Part2(lines)
	fmt.Println(part1)
	fmt.Println(part2)
}

func Part1(lines []string) int {
	var sum int

	for _, i := range lines {
		digits := GetFirstAndLastDigits(i)
		sum += digits
	}

	return sum
}

func Part2(lines []string) int {
	var sum int

	for _, i := range lines {
		digits := GetFirstAndLastDigits(ConvertWrittenNumber(i))
		log.Println("Got:", digits)
		sum += digits
	}

	return sum
}

func GetFirstAndLastDigits(s string) int {
	f, err := FindFirstDigit(s)
	if err != nil {
		log.Fatal("No number detected")
	}
	b, _ := FindFirstDigit(StringReverse(s))
	digits := fmt.Sprintf("%v%v", f, b)
	digitInts, err := strconv.Atoi(digits)
	if err != nil {
		log.Fatal("Cannot convert digits to int")
	}
	return digitInts
}

func FindFirstDigit(s string) (int, error) {
	for _, b := range s {

		if unicode.IsNumber(b) {
			return int(b - '0'), nil
		}
	}
	return 0, fmt.Errorf("no number in string")
}

func StringReverse(s string) string {
	forward := []byte(s)
	size := len(forward)
	backward := make([]byte, size)
	for i := 0; i < size; i++ {
		backward[size-i-1] = forward[i]
	}
	return string(backward)
}

// Might need to split at the first and last number and check if there is a word before the first or after the last
func ConvertWrittenNumber(oldString string) string {
	// var newString string
	nums := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	log.Println("Converting", oldString)
	for s, i := range nums {
		if strings.Contains(oldString, s) {
			oldString = strings.ReplaceAll(oldString, s, fmt.Sprintf("%v%s%v", string(s[0]), i, string(s[len(s)-1])))
		}
	}
	log.Println("New number:", oldString)
	return oldString
}

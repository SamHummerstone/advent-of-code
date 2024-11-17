package stars

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Trebuchet(file string) int {
	var sum int

	input, err := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	if err != nil {
		log.Fatal("Could not open input file:", err)
	}

	for _, i := range lines {
		f, err := FindFirstDigit(i)
		if err != nil {
			log.Println("No number detected")
			continue
		}
		b, _ := FindFirstDigit(StringReverse(i))
		digits, err := strconv.Atoi(string(f) + string(b))
		if err != nil {
			log.Fatal("Cannot convert digits to int")
		}
		sum += digits
	}

	return sum
}

func FindFirstDigit(s string) (int, error) {
	for _, b := range s {

		if unicode.IsNumber(b) {
			return int(b), nil
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

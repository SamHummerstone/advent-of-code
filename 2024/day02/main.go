package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile embed.FS

func main() {
	inputString, err := inputFile.ReadFile("input.txt")
	if err != nil {
		panic("Could not load input file")
	}
	fmt.Println(Part1(string(inputString)))
}

func Part1(reports string) int {
	var count int
	reportSplit := strings.Split(reports, "\n")

	for _, r := range reportSplit {
		if IsReportSafe(ReportStringToInts(r)) {
			count++
		}
	}
	return count
}

func IsReportSafe(report []int) bool {
	var isAscending bool
	var dif, lastNum int
	if firstDif := report[0] - report[1]; firstDif == 0 {
		return false
	} else if firstDif > 0 {
		isAscending = false
	} else {
		isAscending = true
	}

	for i, n := range report {
		if i == 0 {
			lastNum = n
			continue
		}
		dif = lastNum - n
		if dif == 0 {
			return false
		}
		if !isAscending && (dif > 3 || dif < 0) {
			return false
		}
		if isAscending && (dif < -3 || dif > 0) {
			return false
		}
		lastNum = n
	}

	return true
}

func ReportStringToInts(r string) []int {
	var ints []int
	rSplit := strings.Split(r, " ")
	for _, level := range rSplit {
		l, err := strconv.Atoi(level)
		if err != nil {
			panic(fmt.Sprintln("ReportStringToInts:", err))
		}
		ints = append(ints, l)
	}
	return ints
}

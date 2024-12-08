package main

import (
	"regexp"
	"strconv"
)

type Mul struct {
	Left  int
	Right int
}

func (m Mul) Equals() int {
	return m.Left * m.Right
}

func New(s string) (Mul, error) {
	regex := mulExpr
	re := regexp.MustCompile(regex)

	matches := re.FindStringSubmatch(s)

	l, err := strconv.Atoi(matches[1])
	if err != nil {
		return Mul{}, err
	}
	r, err := strconv.Atoi(matches[2])
	if err != nil {
		return Mul{}, err
	}
	newMul := Mul{
		Left:  l,
		Right: r,
	}

	return newMul, nil
}

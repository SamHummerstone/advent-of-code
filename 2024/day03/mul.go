package main

type Mul struct {
	Left  int
	Right int
}

func (m Mul) Equals() int {
	return m.Left * m.Right
}

func New(s string) Mul {
	var newMul Mul

	return newMul
}

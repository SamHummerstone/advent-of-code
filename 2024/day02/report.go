package main

type Report struct {
	Content []int
	Safe    bool
}

// Stolen from michael.schuett at https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// func RemoveIndex(s []int, index int) []int {
// 	ret := make([]int, 0)
// 	ret = append(ret, s[:index]...)
// 	return append(ret, s[index+1:]...)
// }

func (r Report) Dampen(index int) *Report {
	dr := make([]int, 0)
	dr = append(dr, r.Content[:index]...)
	return New(append(dr, r.Content[index+1:]...))
}

func New(rs []int) *Report {
	report := Report{
		Content: rs,
	}
	report.Safe = report.IsSafe()

	return &report
}

func (r *Report) IsSafe() bool {
	var isAscending bool
	var dif, lastNum int
	if firstDif := r.Content[0] - r.Content[1]; firstDif == 0 {
		return false
	} else if firstDif > 0 {
		isAscending = false
	} else {
		isAscending = true
	}

	for i, n := range r.Content {
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

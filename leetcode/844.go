package main

import "fmt"

func main() {
	a := "y#fo##f"
	b := "y#f#o##f"
	fmt.Println(backspaceCompare(a,b))
}

func backspaceCompare(S string, T string) bool {
	var s = make([]rune, 0)
	for _, i := range S {
		if i == '#' {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		} else {
			s = append(s, i)
		}
	}

	var t = make([]rune, 0)
	for _, j := range T {
		if j == '#' {
			if len(t) > 0 {
				t = t[:len(t)-1]
			}
		} else {
			t = append(t, j)
		}

	}
	return string(s) == string(t)
}
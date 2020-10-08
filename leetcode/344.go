package main

import "fmt"

func main() {
	s := "hello"
	b := []byte(s)
	reverseString(b)
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	fmt.Println(s)
}

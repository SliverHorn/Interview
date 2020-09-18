package main

import "fmt"

func ForError() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func ForCorrect() {
	slice := []int{0, 1, 2, 3, 4}
	m := make(map[int]*int)
	for key, val := range slice {
		value := val
		m[key] = &value
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func main() {
	ForError()
	ForCorrect()
}

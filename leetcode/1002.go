package main

import (
	"fmt"
	"math"
)

func main() {
	A := []string{"bella", "label", "roller"}
	B := commonChars(A)
	fmt.Println(B)
}

func commonChars(A []string) []string {
	if A == nil || len(A) == 0 {
		return []string{}
	}
	alphabet := [26]int{}
	for index := range alphabet {
		alphabet[index] = math.MaxUint16
	}
	alphabetInWork := [26]int{}
	for _, word := range A {
		for _, char := range []byte(word) { // 编译器技巧-在这里我们不会分配新的内存
			alphabetInWork[char-'a']++
		}

		for i := 0; i < 26; i++ {
			if alphabetInWork[i] < alphabet[i] { // 缩小频次，使得统计的公共频次更加准确
				alphabet[i] = alphabetInWork[i]
			}
		}

		for i := range alphabetInWork { // 重置状态
			alphabetInWork[i] = 0
		}
	}
	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		for g := 0; g < alphabet[i]; g++ {
			result = append(result, string(i+'a'))
		}
	}
	return result
}

package main

import "fmt"

func SliceError() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func SliceCorrect() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
}

//func funcMui(x,y int)(sum int,error){
//	return x+y,nil
//}

func funcMui1(x,y int)(int,error){
	return x+y,nil
}

func funcMui2(x,y int)(sum int,err error){
	return x+y,nil
}

func main() {
	SliceError()
	SliceCorrect()
	fmt.Println(funcMui1(1, 2))
	fmt.Println(funcMui2(1, 2))
}


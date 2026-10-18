package main

import "fmt"

var b []int = make([]int, 3)

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	a = append(a, 4)
	fmt.Println(a)
	fmt.Println(cap(a))
	a = append(a[:3], a[4:]...)
	fmt.Println(a)
	modify()
}

func modify() {
	b[0] = 999
	fmt.Println(b)
}

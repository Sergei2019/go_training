package main

import "fmt"

func main() {
	// var k = []int{5, 6, 7}
	var m []int
	var s int
	s = sum_slice(m)
	fmt.Println(s)
}

func sum_slice(a []int) int {
	var s int
	for _, v := range a {
		s = s + v
	}
	return s
}

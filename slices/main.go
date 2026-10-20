package main

import "fmt"

func main() {
	var s = []int{}

	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}
	fmt.Println(len(s))
	fmt.Println(cap(s))
	var a = make([]int, len(s))
	copy(a, s)
	fmt.Println(a)

}

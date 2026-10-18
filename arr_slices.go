package main

import "fmt"

func main() {
	var slices = [...]string{
		"test",
		"test2",
		"test3",
	}
	fmt.Println(slices[0:2])

	var slice_two = make([]int, 0, 10)

	for i := 0; i <= 10; i++ {
		slice_two = append(slice_two, i)
	}
	fmt.Println(slice_two)
}

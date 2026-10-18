package main

import "fmt"

func main() {
	var a []int64
	var b = []int64{1, 2, 3, 4}
	for _, v := range b {
		a = append(a, v)
	}
	fmt.Println(a)

}

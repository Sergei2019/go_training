package main

import "fmt"

var list_ []int

func main() {
	sliceTricks()
	fmt.Println(list_)
}

func sliceTricks() {
	list_ = append(list_, 1)
}

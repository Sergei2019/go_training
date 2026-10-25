package main

import "fmt"

func main() {

	s := make(map[string]int)
	s["a"] = 1

	fmt.Println(s["b"])
}

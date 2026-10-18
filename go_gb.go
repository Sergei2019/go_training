package main

import "fmt"

func main() {
	size := 4831838208.0 //bites
	kb := size / 1024.0 / 1024.0 / 1024.0

	fmt.Print(kb)
}

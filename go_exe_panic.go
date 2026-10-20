package main

import "fmt"

func main() {
	var a, b = 3, 0
	fmt.Println(div(a, b))
}

func div(a int, b int) int {
	return a / b
}

func first_letter(s string) string {
	return string(s[0])
}

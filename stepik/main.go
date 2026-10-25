package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	num := fmt.Sprintf("%d", a)
	fmt.Printf("%c", num[len(num)-2])

}

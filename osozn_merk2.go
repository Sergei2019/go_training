package main

import "fmt"

func main() {
	// var rating float64 = 8.4
	// var title string = "film"
	// var year int = 2026

	second_name := "film3"
	fmt.Println(second_name)
	second_name = "6"
	fmt.Println(second_name)
	var rating float64
	fmt.Println(rating)
	const max_rating = 10.0
	fmt.Println(max_rating)
	const (
		pi = 3.14
		g  = 9.8
	)
	fmt.Println(pi)
	var zeroString string
	fmt.Println(zeroString)
	var zeroRune rune = 'a'
	fmt.Printf("Zero runa ravna %c\n", zeroRune)

}

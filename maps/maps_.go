package main

import "fmt"

func main() {
	fmt.Println(countChars("aaashshs"))
}

func countChars(a string) map[rune]int {
	mapa := make(map[rune]int, len(a))
	r := []rune(a)
	for i := 0; i < len(a); i++ {
		key := r[i]
		mapa[key] = mapa[key] + 1
	}
	return mapa

}

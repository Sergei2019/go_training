package main

import "fmt"

func main() {
	fmt.Println(hasDuplicates([]int{5, 5, 7}))
}

func hasDuplicates(a []int) bool {
	mapa := make(map[int]int)
	for _, i := range a {

		if _, ok := mapa[i]; ok {
			return true
		}
		mapa[i] = i
	}
	return false

}

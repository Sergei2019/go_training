package main

import (
	"fmt"
	"slices"
)

func main() {
	var s = []string{"banana", "cocoa", "apple", "banana", "apple"}
	s = remove_dubl(s)
	fmt.Println(s)

}

func remove_dubl(s []string) []string {
	var k []string
	for _, v := range s {
		if !slices.Contains(k, v) {
			k = append(k, v)
		}
	}
	return k
}

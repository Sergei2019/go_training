package main

import (
	"fmt"
	"time"
)

// var x, y, z int = 4, 5, 6

func main() {
	// num := 3
	// fmt.Println(x, y, z, num)
	// fmt.Println(true && false)
	// fmt.Println(true || true)
	// fmt.Println("5.0/2.0", 5.0/2.0)
	m := dur()
	fmt.Println(m.Minutes())

}

func dur() time.Duration {
	m, _ := time.ParseDuration("1h30m")
	return m
}

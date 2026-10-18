package main

import "fmt"

func main() {
	duration := 8520
	minutes := duration / 60
	hours := minutes / 60
	real_min := minutes % 60
	seconds := real_min / 60

	fmt.Printf("Длитеьность фильма составляет %d часов %d минут %d секунд", hours, real_min, seconds)
}

package main

import "fmt"

func main() {
	s := []float64{7.3, 4.5, 2.0, 10.0, 3.4}
	min_rat := 5.0
	fmt.Println(filter_by(s, min_rat))
}

func filter_by(s []float64, min_r float64) []float64 {
	var a = make([]float64, 0, len(s))
	for _, v := range s {
		if v >= min_r {
			a = append(a, v)
		}
	}
	return a
}

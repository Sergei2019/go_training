package main

import (
	"fmt"
	"math"
	"errors"
)

const pi float64 = 3.14
var name string = "Круг"


func main(){
	CircleS(10.5)
	DlinaKruga(5)
}

func CircleS(radius float64){
	var s float64
	if radius < 0 {
		fmt.Println(errors.New("ошибка , радиус меньше 0"))
	} else {
		s = pi * math.Pow(radius,2)
		fmt.Println(s)
		fmt.Println(name)
	}
}


func DlinaKruga(R float64){
	var D float64
	D = 2*pi*R
	fmt.Printf("Длина %s \n", name)
	fmt.Println(D)
}
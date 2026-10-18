package main

import (
	"fmt"
	"errors"
)


var a = []int{1,2,3,4}

func main() {
	fmt.Println(add(1,2))
	fmt.Println(divide(1,0))
	fmt.Println(divide(1,2))
	loop_zaloop()
}

func add(a int, b int) int{
	return a + b
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	} else {
		return a / b, nil
	}
}

func  loop_zaloop()  {
	for i:=0; i < 10; i++ {
		fmt.Println("плюс один")
	}
	
}

func  loop_zaloop2()  {
	for i:=0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
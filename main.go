package main

import (
	"fmt"
)

// var для глобальных переменных, а := для функций
var message = "тест"
var b string = "тест2"
var mem *string = &message // указатель - адрес в ячейке памяти
var chislo = 25
var pt_chisla *int = &chislo

func subMessage(alertMessage string) {
	fmt.Printf("Сервис запущен с сообщением %s \n", alertMessage)
}

func main() {
	var a = 1
	subMessage("Привет")
	c := "тест3"
	var mem = "тест4"
	fmt.Println("Hello")
	fmt.Println(mem)
	//fmt.Println(*mem) //получаем значение
	fmt.Println(pt_chisla)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("Число %d равно %s", 2, "двум\n")
	switch {
	case a > 0:
		fmt.Println("good")
	default:
		fmt.Println("bad")
	}

}

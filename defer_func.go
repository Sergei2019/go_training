package main

import "fmt"

// стек отложенных вызвовов. То есть каждая функция кладется в стек. Первый вошел-последнийвышел
func main() {
	defer func() {
		fmt.Println("Я анонимгная функция")
	}()
	defer func() {
		fmt.Println("Я анонимгная функция2")
	}()

	fmt.Println("прривет")
}

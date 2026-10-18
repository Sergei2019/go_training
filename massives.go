package main

import "fmt"

var list_ = []int{1,2,3}
var mapa = map[string]int {
	"user1": 1,
	"user2": 2,
	"user3": 3,
}

type User struct {
	user string
	age int
	sex string
	status bool
}

func main() {
	// list_modifier(list_)
	// list_print()
	// list_modifier(list_)
	// list_modifier2()
	mapa_iter()
	fmt.Println(struct_init())
}

func list_modifier(lst []int) {
	for i := 0; i < len(lst); i++ {
		fmt.Println(i)
	}
}

func list_print() {
	fmt.Println(list_[1])
	list_ = append(list_, 5)
}


func list_modifier2() {
	for _, i := range list_ {
		fmt.Println(i)
	}
}


func mapa_iter() {
	for _, v := range mapa {
		fmt.Println(v)
	}
}


func struct_init() User{
	return User{"Вася", 40, "м", true}
}
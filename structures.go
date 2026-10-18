package main

import "fmt"

func main() {
	user1 := create_user(1, "Вася", 33, true)
	fmt.Printf("Пользователь %s создан\n", user1.name)

}

type User struct {
	id     int
	name   string
	age    int
	status bool
}

func create_user(i int, name string, age int, status bool) User {
	user := User{i, name, age, status}
	return user
}

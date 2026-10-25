package main

import "fmt"

func main() {
	type User struct {
		name string
		age  int
		sex  string
	}
	var u User
	_, err := fmt.Scan(&u.name, &u.age, &u.sex)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)

}

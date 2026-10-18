package main

import "fmt"


func main(){
	var todo_list = [3]string{
		"тест1",
		"тест2",
		"тест3",
	}
	fmt.Println(todo_list)

	for index, val := range todo_list {
		fmt.Printf("%d. %s\n",index, val)
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// for i := 0; ; i++ {
	// 	fmt.Println(i)
	// }
	fmt.Println(todo_list[0:2])
}

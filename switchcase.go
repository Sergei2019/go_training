package main

import "fmt"

func main() {
	var lst_of_commands = []string{"play", "stop", "remove"}
	for _, v := range lst_of_commands {
		switch {
		case v == "play":
			fmt.Println(v)
		case v == "remove":
			fmt.Println("bad")

		}
	}

}

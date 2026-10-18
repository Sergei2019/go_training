package main

type employee struct {
	id     int
	name   string
	age    int
	status bool
}

type storage interface {
	insert(e employee)
	delete(id int)
}

func main() {

}

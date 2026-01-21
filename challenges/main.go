package main

import "fmt"

func main() {

	type People struct {
		name string
		age  int
	}
	kevin := People{
		name: "kevin",
		age:  30,
	}
	fmt.Println(kevin)

}

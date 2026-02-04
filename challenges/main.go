package main

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	ID    int    `json: "id`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	t := Task{
		ID:    1,
		Title: "Buy milk",
		Done:  false,
	}
	// struct ➡️ json
	data, _ := json.Marshal(t)
	fmt.Println(data)

}

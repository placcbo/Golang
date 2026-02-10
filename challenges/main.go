package main

import (
	"fmt"

	"time"
)

func main() {
	sendEmail("kevn")
	fmt.Scanln()

}

func sendEmail(email string) {
	func() {
		time.Sleep(10 * time.Second)
		fmt.Printf("hello %s ", email)
	}()
}

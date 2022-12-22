package main

import (
	"fmt"
	"hello-world/greeting"
)

var greet string

func main() {
	greet = greeting.HelloWorld()

	fmt.Println(greet)
}

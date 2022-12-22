package main

import "fmt"

// HelloWorld greets the world.
func HelloWorld() string {
	return "Hello, World!"
}

var greeting string

func main() {
	greeting = HelloWorld()

	fmt.Println(greeting)
}

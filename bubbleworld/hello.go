package main

import "fmt"

func Hello(worldName string) string {
	return "Hello, " + worldName + " World."
}

func main() {
	fmt.Println(Hello("some"))
}

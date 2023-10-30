package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloSuffix = " World."

func Hello(worldName string) string {
	return englishHelloPrefix + worldName + englishHelloSuffix
}

func main() {
	fmt.Println(Hello("some"))
}

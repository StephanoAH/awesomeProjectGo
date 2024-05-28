package main

import "fmt"

const helloPrefix string = "Hello, "

func Hello(param string) string {
	if param == "" {
		param = "World"
	}
	return helloPrefix + param
}

func main() {
	fmt.Println(Hello("Stephano"))
}

package main

import "fmt"

func Hello(param string) string {
	return "Hello, " + param
}

func main() {
	fmt.Println(Hello("Stephano"))
}

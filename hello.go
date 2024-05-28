package main

import "fmt"

const helloPrefix string = "Hello, "

func Hello(param string) string {
	return helloPrefix + param
}

func main() {
	fmt.Println(Hello("Stephano"))
}

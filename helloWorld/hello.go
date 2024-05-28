package main

import (
	"fmt"
)

type Params struct {
	subject, lang string
}

var helloPrefix string = "Hello, "

func GreetingLangs(lang string) string {
	switch lang {
	case "Spanish":
		helloPrefix = "Hola, "
	case "Korean":
		helloPrefix = "안녕하세요, "
	case "Italian":
		helloPrefix = "Ciao, "
	default:
		helloPrefix = "Hello, "
	}
	return helloPrefix
}

func Hello(params Params) string {
	if params.subject == "" {
		params.subject = "World"
	}
	GreetingLangs(params.lang)
	return helloPrefix + params.subject
}

func main() {
	fmt.Println(Hello(Params{subject: "World", lang: "Italian"}))
}

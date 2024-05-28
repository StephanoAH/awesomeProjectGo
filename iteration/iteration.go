package iteration

import "fmt"

type Params struct {
	element string
	times   int
}

func Iteration(params Params) string {
	var iteratedString string
	for i := 0; i < params.times; i++ {
		iteratedString += params.element
	}
	return iteratedString
}

func iteration() {
	fmt.Println(Iteration(Params{element: "h", times: 5}))
}

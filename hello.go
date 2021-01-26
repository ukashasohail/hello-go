package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	var returned int = max(5, 10)

	fmt.Println(returned)
}

func max(num1, num2 int) int {

	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

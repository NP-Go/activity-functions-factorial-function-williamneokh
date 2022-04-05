package main

import (
	"fmt"
	"strconv"
)

//Declare factorial function here
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	//Insert Code here
	var factorialVal string
	fmt.Println("Input a value")
	fmt.Scanln(&factorialVal)

	factorialValue, _ := strconv.ParseInt(factorialVal, 10, 0)

	result := factorial(int(factorialValue))
	fmt.Println(result)

}

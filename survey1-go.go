package main

import "fmt"

func main() {
	factorialMap := make(map[int]int)
	fmt.Println(factorial(7, factorialMap))
	fmt.Println(factorialMap)
}

func factorial(n int, factorialMap map[int]int) int {

	if n <= 1 {
		return 1
	} else if factorialMap[n] != 0 {
		return factorialMap[n]
	} else {
		var newFactorial int = n * factorial(n-1, factorialMap)
		factorialMap[n] = newFactorial
		return newFactorial
	}

}

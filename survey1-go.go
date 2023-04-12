//Co-authors: Kyle Kuminkoski and Grant Stumpf

package main

import "fmt"

func main() {
	var n int
	var k int
	factorialMap := make(map[int]int)
	fmt.Print("Enter a value for n: ")
	fmt.Scan(&n)
	fmt.Print("Enter a value for k: ")
	fmt.Scan(&k)
	fmt.Print("Hypercake value: ")
	fmt.Println(hypercake(n, k, factorialMap))
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

func combinations(n int, r int, factorialMap map[int]int) int {
	if 0 <= r && r >= n {
		return 0
	} else {
		return factorial(n, factorialMap) / (factorial(r, factorialMap) * factorial(n-r, factorialMap))
	}
}

func hypercake(n int, k int, factorialMap map[int]int) int {
	var value = 0
	for k >= 0 {
		value = value + combinations(n, k, factorialMap)
		k = k - 1
	}
	return value
}

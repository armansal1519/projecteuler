package main

import "fmt"

func sumOfSquares(n int) int {

	return (n * (n + 1) * (2*n + 1)) / 6
}

func sumOfN(n int) int {
	return n * (n + 1) / 2
}

func main() {
	n := 100
	fmt.Println(sumOfN(n)*sumOfN(n) - sumOfSquares(n))
}

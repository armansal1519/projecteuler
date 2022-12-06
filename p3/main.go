package main

import (
	"fmt"
	"math"
)

func main() {

	n := 600851475143
	f := n
	factors := make(map[int]int, 0)
	fmt.Println(math.Sqrt(float64(n)) + 1)
	for i := 2; i < int(math.Sqrt(float64(f))+1); i++ {
		for n%i == 0 {
			factors[i] = 0
			n = n / i
		}
	}
	fmt.Println(factors)
}

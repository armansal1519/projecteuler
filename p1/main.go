package main

import "fmt"

func main() {
	sum := 0
	n := 1000

	for i := 0; i < n; i += 3 {
		sum += i
	}
	for i := 0; i < n; i += 5 {
		if i%15 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}

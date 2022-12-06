package main

import "fmt"

func main() {
	fibList := []int{1, 2}
	for fibList[len(fibList)-1] < 4000000 {
		l := len(fibList)
		fibList = append(fibList, fibList[l-1]+fibList[l-2])
	}
	sum := 0
	for i := 1; i < len(fibList); i += 3 {
		sum += fibList[i]
	}
	fmt.Println(sum)
}

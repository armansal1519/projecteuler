package main

import (
	"fmt"
	"sort"
)

func isPalindrome(n int) bool {
	var reminder int
	arr := make([]int, 0)
	for n > 0 {
		reminder = n % 10
		arr = append(arr, reminder)
		n /= 10
	}
	for i := 0; i < (len(arr) / 2); i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}

func palindrome(a int, b int) bool {
	if isPalindrome(a * b) {
		return true
	}
	return false
}

func main() {
	array := make([]int, 0)
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if palindrome(i, j) {
				array = append(array, i*j)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	fmt.Println(array[0])
}

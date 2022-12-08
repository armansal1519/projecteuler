package main

import "fmt"

func isDivz(n int) bool {
	a := make([]int, 0)
	for i := 1; i < 21; i++ {
		a = append(a, i)
	}
	for _, i := range a {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {

	for i := 1; true; i++ {
		if isDivz(i) {
			fmt.Println(i)
			break
		}
	}

}

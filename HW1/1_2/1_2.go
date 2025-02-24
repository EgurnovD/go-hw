package main

import "fmt"

func Solution_1_2(slice []int) []int {
	for i := range slice {
		p := &slice[i]
		*p += 5
	}
	return slice
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl := Solution_1_2(numbers)

	fmt.Println(sl)
}

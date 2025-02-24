package main

import "fmt"

func Solution_1_1(slice []int) ([]int, int) {
	sum := 0
	for i := range slice {
		if slice[i]%2 == 0 {
			slice[i] = 1
		}
		sum += slice[i]
	}
	return slice, sum
}

func main() {
	numbers := []int{3, 5, 7, 2, 7, 8, 6, 4, 7, 0, 1, 7, 4, 8, 10, 3, 6, 8, 5, 4, 12, 3}
	sl, sum := Solution_1_1(numbers)

	fmt.Println("Sum = ", sum)
	fmt.Println(sl)
}

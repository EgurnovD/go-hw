package main

import "fmt"

func Solution_1_3(slice []int) (int, int) {
	min, max := slice[0], slice[0] // assume slice not empty
	for _, v := range slice {
		if v%2 == 0 {
			switch {
			case v < min:
				min = v
			case v > max:
				max = v
			}
		}
	}
	return min, max
}

func main() {
	numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}
	min, max := Solution_1_3(numbers)

	fmt.Printf("Min = %v, Max = %v", min, max)
}

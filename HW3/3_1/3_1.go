package main

import (
	"fmt"
)

func main() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40} {
		fmt.Printf("%v %d\n", &value, value)
		numbers = append(numbers, &value)
	}
	fmt.Println("----------")
	for _, number := range numbers {
		fmt.Printf("%v %d\n", number, *number)
	}
}

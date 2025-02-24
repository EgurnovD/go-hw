package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func iterate(name string, n int) {
	defer wg.Done()

	for i := 0; i <= n; i++ {
		// fmt.Println(name, i)
		fmt.Print(i)
	}
	fmt.Println(' ', name, "finished")
}
func main() {

	runtime.GOMAXPROCS(1)
	wg.Add(3)
	go iterate("a", 1)
	go iterate("b", 2)
	go iterate("c", 3)
	wg.Wait()
}

// Output
// 012345678910
// 012345678910
// 012345678910

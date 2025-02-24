package main

import (
	"fmt"
	"runtime"
	"sync"
)

func process(n_procs int, n int) {
	runtime.GOMAXPROCS(n_procs)
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
	wg := sync.WaitGroup{}
	wg.Add(n)
	fmt.Println("  Before loop")
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("      Почему, КОЛЯ?", i)
		}(i)
	}
	fmt.Println("  After loop")
	wg.Wait()
}

func main() {
	// About Go Scheduler
	// [1] https://backendinterview.ru/goLang/scheduler.html
	// [2] https://habr.com/ru/articles/858490/
	process(0, 5)
	process(1, 5)
	fmt.Println("Паника")
}

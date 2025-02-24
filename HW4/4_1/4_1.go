package main

import (
	"fmt"
	"strings"
)

func words_counter(ch_in <-chan string, ch_out chan<- int, n int) {
	fmt.Println("In goroutine ", n)
	for s := range ch_in {
		fmt.Println("processing in ", n)
		ch_out <- len(strings.Fields(s))
	}
}

func main() {
	n_routines := 2
	lines := []string{"Всем привет!",
		"Следующая лекция в среду",
		"Увидимся на лекции!"}
	ch_in := make(chan string)
	ch_out := make(chan int, len(lines))
	// создаем несколько исполнителей
	for i := 0; i < n_routines; i++ {
		go words_counter(ch_in, ch_out, i)
	}
	// отправляем строки в канал
	for _, line := range lines {
		fmt.Println("emit ", line)
		ch_in <- line
	}
	close(ch_in)
	// забираем результаты
	fmt.Println("Results incoming")
	for i := 0; i < len(lines); i++ {
		res := <-ch_out
		fmt.Println("Word count: ", res)
	}
}

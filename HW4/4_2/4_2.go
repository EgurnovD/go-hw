package main

import (
	"fmt"
	"time"
)

func dummy_user(id int, num_messages int, ch_out chan<- string) {
	for i := 0; i < num_messages; i++ {
		ch_out <- fmt.Sprintf("[User %v] Message %v", id, i)
	}
}

func main() {
	num_users := 3
	num_messages := 5
	chat := make(chan string)
	// создаем несколько пользователей
	for u := 0; u < num_users; u++ {
		go dummy_user(u, num_messages, chat)
	}
	// выводим чат
	for {
		select {
		case message := <-chat:
			fmt.Println(message)
		default:
			time.Sleep(time.Second)
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func tryTest() func() {
	fmt.Println("tryTest")
	return func() {
		fmt.Println("tryTest2")
	}
}

func Option1() {
	defer fmt.Println("Первое время:", time.Now())
	defer tryTest()()
	time.Sleep(2 * time.Second)
	defer fmt.Println("Второе время", time.Now())
	fmt.Println("-----")
}
func Option2() {
	defer fmt.Println("Первое время:", time.Now())
	defer tryTest()
	time.Sleep(2 * time.Second)
	defer fmt.Println("Второе время", time.Now())
	fmt.Println("-----")
}

func main() {
	// В первом случае tryTest выполнится в обычном порядке, а возвращенная
	// анонимная функция ляжет в стэк и будет выполнена при разворачивании defer
	Option1()
	fmt.Println()
	// Во вотром случае отложится исполнение самого tryTest.
	// Функция будет выполнена при обратном проходе по defer.
	// Возвращенная анонимная функция не вызывается
	Option2()
}

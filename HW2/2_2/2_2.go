package main

import "fmt"

type Balance map[string]float64

func AddCategory(balance Balance, name string) {
	if _, exists := balance[name]; !exists {
		balance[name] = 0
	}
}
func AddExpense(balance Balance, name string, amount float64) {
	balance[name] += amount
}
func PrintTotal(balance Balance) {
	var sum float64
	for _, amount := range balance {
		sum += amount
	}
	fmt.Printf("Total: %.2f\n", sum)
}
func PrintState(balance Balance) {
	PrintTotal(balance)
	for name, amount := range balance {
		fmt.Printf("  %v: %.2f\n", name, amount)
	}
	fmt.Println()
}

func main() {
	balance := make(Balance)
	PrintState(balance)
	AddCategory(balance, "Food")
	PrintState(balance)
	AddExpense(balance, "Food", 10)
	PrintState(balance)
	AddExpense(balance, "Food", 5)
	AddExpense(balance, "Fun", 7)
	PrintState(balance)
}

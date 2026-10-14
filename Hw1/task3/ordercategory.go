package main

import "fmt"

func main() {
	var amount int64
	fmt.Print("сумма заказа: ")
	fmt.Scan(&amount)

	if amount < 0 {
		fmt.Println("ошибка, отрицательная сумма заказа")
		return
	}

	switch {
	case amount < 10_000:
		fmt.Println("small")
	case amount >= 10_000 && amount <= 49_999:
		fmt.Println("regular")
	default:
		fmt.Println("large")
	}
}

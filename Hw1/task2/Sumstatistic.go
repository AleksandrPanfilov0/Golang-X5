package main

import "fmt"

func main() {
	var count int
	fmt.Print("количество заказов: ")
	fmt.Scan(&count)
	amounts := make([]int64, count)

	for i := 0; i < count; i++ {
		fmt.Printf("заказ %d: ", i+1)
		fmt.Scan(&amounts[i])
	}

	if len(amounts) == 0 {
		fmt.Println("пусто")
		return
	}
	min := amounts[0]
	max := amounts[0]
	var sum int64

	for _, x := range amounts {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
		sum += x
	}

	mean := float64(sum) / float64(len(amounts))
	fmt.Println("минимум:", min)
	fmt.Println("максимум:", max)
	fmt.Println("сумма:  ", sum)
	fmt.Printf("среднее: %.2f\n", mean)
}

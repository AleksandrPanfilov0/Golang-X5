package main

import (
	"errors"
	"fmt"
)

type Item struct {
	Name         string
	PriceKopecks int64
}

type Order struct {
	ID     int
	Items  []Item
	Status string
}

func newOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, errors.New("нужен положительный id")
	}
	return Order{ID: id, Status: "new"}, nil
}

func addItem(order *Order, item Item) error {
	if order == nil {
		return errors.New("заказ не задан")
	}
	if item.Name == "" {
		return errors.New("имя не может быть пустым")
	}
	if item.PriceKopecks < 0 {
		return errors.New("цена не может быть отрицательной")
	}
	order.Items = append(order.Items, item)
	return nil
}

func total(order Order) int64 {
	var result int64
	for _, item := range order.Items {
		result += item.PriceKopecks
	}
	return result
}

func main() {
	statusNames := make(map[string]string)
	statusNames["new"] = "новый"
	statusNames["paid"] = "оплачен"
	order, err := newOrder(42)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	_ = addItem(&order, Item{Name: "Кофе", PriceKopecks: 19_990})
	_ = addItem(&order, Item{Name: "Фильтр", PriceKopecks: 2_500})

	err = addItem(&order, Item{Name: "", PriceKopecks: 1_000})
	if err != nil {
		fmt.Println("ошибка добавления", err)
	}
	fmt.Printf("заказ %d, статус %s, сумма %d коп.\n",
		order.ID, statusNames[order.Status], total(order))
}

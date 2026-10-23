package main

import (
	"testing"
)

func TestOrderZero(t *testing.T) {
	order, err := newOrder(48)
	if err != nil {
		t.Errorf("ошибка при создании заказа %v", err)
		return
	}
	result := total(order)
	expect := int64(0)
	if result != expect {
		t.Errorf("ошибка заказ должен иметь сумму %d, получено %d", expect, result)
	}
}

func TestTotal(t *testing.T) {
	order, err := newOrder(48)
	if err != nil {
		t.Errorf("ошибка создания заказа %v", err)
		return
	}
	_ = addItem(&order, Item{Name: "Кофе", PriceKopecks: 19_990})
	_ = addItem(&order, Item{Name: "Фильтр", PriceKopecks: 2_500})
	result := total(order)
	expect := int64(19_990 + 2_500)
	if result != expect {
		t.Errorf("ошибка, ожидалась сумма %d, получено %d", expect, result)
	}
}

func TestAddItem(t *testing.T) {
	order, err := newOrder(48)
	if err != nil {
		t.Errorf("ошибка создания заказа %v", err)
		return
	}

	if err := addItem(nil, Item{Name: "Кофе", PriceKopecks: 100}); err == nil {
		t.Errorf("ожидалась ошибка, передан nil вместо заказа")
	}

	if err := addItem(&order, Item{Name: "", PriceKopecks: 100}); err == nil {
		t.Errorf("пропустили товар с пустым именем")
	}

	if err := addItem(&order, Item{Name: "Кофе", PriceKopecks: -50}); err == nil {
		t.Errorf("пропустили отрицательную цену")
	}
}

func TestOrderCopy(t *testing.T) {
	order, _ := newOrder(1)
	copyOrder := order
	copyOrder.Status = "отменен"
	if order.Status == "отменен" {
		t.Errorf("оригинальный заказ изменился, а не должен был")
	}
}

func TestMapInitialization(t *testing.T) {
	statusNames := make(map[string]string)
	statusNames["new"] = "новый"
	statusNames["paid"] = "оплачен"
	if statusNames["new"] != "новый" {
		t.Errorf("map не была корректно инициализирована")
	}
}

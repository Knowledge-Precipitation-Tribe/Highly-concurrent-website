package repositories

import (
	"Highly-concurrent-website/datamodels"
	"fmt"
	"testing"
)

func TestOrderMangerRepository_Insert(t *testing.T) {
	order := &datamodels.Order{
		UserID:      2,
		ProductId:   3,
		OrderStatus: 4,
	}
	orderManger := &OrderMangerRepository{
		table: "website.order",
	}
	productID, err := orderManger.Insert(order)
	if err != nil {
		panic(err)
	}
	order.ID = productID
}

func TestOrderMangerRepository_SelectByKey(t *testing.T) {
	orderManger := &OrderMangerRepository{
		table: "website.order",
	}
	order, err := orderManger.SelectByKey(274)
	if err == nil {
		fmt.Println(order)
	}
}

func TestOrderMangerRepository_SelectAll(t *testing.T) {
	orderManger := &OrderMangerRepository{
		table: "website.order",
	}
	order, err := orderManger.SelectAll()
	if err == nil {
		fmt.Println(order[0])
		fmt.Println(order[1])
	}
}

func TestOrderMangerRepository_Update(t *testing.T) {
	order := &datamodels.Order{
		ID:          274,
		UserID:      0,
		ProductId:   0,
		OrderStatus: 0,
	}
	orderManger := &OrderMangerRepository{
		table: "website.order",
	}
	err := orderManger.Update(order)
	if err != nil {
		panic(err)
	}
}

func TestOrderMangerRepository_SelectAllWithInfo(t *testing.T) {
	orderManger := &OrderMangerRepository{
		table: "website.order",
	}
	order, err := orderManger.SelectAllWithInfo()
	if err == nil {
		fmt.Println(order)
	} else {
		panic(err)
	}
}

func TestOrderMangerRepository_Delete(t *testing.T) {
	orderManger := &OrderMangerRepository{
		table: "website.order",
	}
	isOK := orderManger.Delete(275)
	fmt.Println(isOK)
}

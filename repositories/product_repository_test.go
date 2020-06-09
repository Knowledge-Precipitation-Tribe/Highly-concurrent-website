package repositories

import (
	"Highly-concurrent-website/datamodels"
	"fmt"
	"testing"
)

var product datamodels.Product

func TestProductManager_Insert(t *testing.T) {
	product = datamodels.Product{
		ProductName:"test1",
		ProductNum:5,
		ProductImage:"test1",
		ProductUrl:"test",
	}
	productManager := &ProductManager{
		table:"product",
	}
	id ,err := productManager.Insert(product)
	if err != nil{
		panic(err)
	}
	product.ID = id
}

func TestProductManager_Update(t *testing.T) {
	product.ProductName = "test1"
	productManager := &ProductManager{
		table:"product",
	}
	err := productManager.Update(product)
	if err != nil{
		panic(err)
	}
}

func TestProductManager_SelectByKey(t *testing.T) {
	productManager := &ProductManager{
		table:"product",
	}
	product, err := productManager.SelectByKey(3)
	if err != nil{
		panic(err)
	}
	fmt.Println(product)
}

func TestProductManager_SelectAll(t *testing.T) {
	productManager := &ProductManager{
		table:"product",
	}
	products, err := productManager.SelectAll()
	if err != nil{
		panic(err)
	}
	fmt.Println(products)
}

func TestProductManager_Delete(t *testing.T) {
	productManager := &ProductManager{
		table:"product",
	}
	productManager.Delete(4)
}
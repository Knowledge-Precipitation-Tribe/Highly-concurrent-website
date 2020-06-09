package repositories

import (
	"Highly-concurrent-website/datamodels"
	"database/sql"
)

type IProduct interface {
	Conn() error
	Insert(product *datamodels.Product)(int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64)(*datamodels.Product, error)
	SelectAll()([]*datamodels.Product, error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct{
	return &ProductManager{table:table, mysqlConn:db}
}


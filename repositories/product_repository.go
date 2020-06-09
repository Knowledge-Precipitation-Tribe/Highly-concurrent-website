package repositories

import (
	"Highly-concurrent-website/common"
	"Highly-concurrent-website/datamodels"
	"database/sql"
	"strconv"
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

//func NewProductManager(table string, db *sql.DB) IProduct{
//	return &ProductManager{table:table, mysqlConn:db}
//}

//创建数据库连接
func (p *ProductManager) Conn() error{
	if p.mysqlConn == nil{
		mysql, err := common.NewMysqlConn()
		if err != nil{
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == ""{
		p.table = "product"
	}
	return nil
}

//插入商品数据
func (p *ProductManager) Insert(product datamodels.Product) (int64, error){
	if err := p.Conn(); err != nil{
		return 0, nil
	}
	sql := "INSERT INTO "+ p.table +" (productName, productNum, " +
		"productImage, productUrl) VALUES (?,?,?,?)"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil{
		return 0, err
	}
	result, err := stmt.Exec(product.ProductName,
		product.ProductNum,
		product.ProductImage,
		product.ProductUrl)
	if err != nil{
		return 0, err
	}
	return result.LastInsertId()
}

//删除商品数据
func (p *ProductManager) Delete(productID int64) bool{
	if err := p.Conn(); err != nil{
		return false
	}
	sql := "DELETE FROM "+ p.table +" WHERE ID = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil{
		return false
	}
	_, err = stmt.Exec(productID)
	if err != nil{
		return false
	}
	return true
}

//更新商品信息
func (p *ProductManager)Update(product datamodels.Product) error{
	if err := p.Conn(); err != nil{
		return err
	}

	sql := "UPDATE " + p.table +" SET productName = ?, " +
		"productNum = ?, productImage=?, productUrl=? " +
		"WHERE ID = "+strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil{
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum,
		product.ProductImage, product.ProductUrl)
	if err != nil{
		return err
	}
	return nil
}

//根据ID查询商品信息
func (p *ProductManager) SelectByKey(productID int64)(*datamodels.Product, error){
	if err := p.Conn(); err != nil{
		return &datamodels.Product{}, err
	}
	sql := "SELECT * FROM " + p.table +" WHERE ID = " + strconv.FormatInt(productID, 10)
	row, err := p.mysqlConn.Query(sql)
	if err != nil{
		return &datamodels.Product{}, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0{
		return &datamodels.Product{}, err
	}
	productResult := &datamodels.Product{}
	common.DataToStructByTagSql(result, productResult)
	return productResult, nil
}
package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" consite:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" consite:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" consite:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" consite:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" consite:"ProductUrl"`
}

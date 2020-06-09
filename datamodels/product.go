package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"id" conSite:"id"`
	ProductName  string `json:"ProductName" sql:"productName" conSite:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" conSite:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" conSite:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" conSite:"ProductUrl"`
}

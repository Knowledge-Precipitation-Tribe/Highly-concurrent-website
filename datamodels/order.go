package datamodels

type Order struct{
	ID int64 `sql:"ID"`
	UserID int64 `sql:"userID"`
	ProductId int64 `sql:"productID"`
	OrderStatus int64 `sql:"orderStatus"`
}

const (
	OrderWait = iota
	orderSuccess //自增为1
	OrderFailed //自增为2
)
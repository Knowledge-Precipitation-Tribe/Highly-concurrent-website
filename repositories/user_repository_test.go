package repositories

import (
	"Highly-concurrent-website/datamodels"
	"testing"
)

func TestUserManagerRepository_Insert(t *testing.T) {
	user := &datamodels.User{
		NickName:"tom",
		UserName:"Jack",
		HashPassword:"2222",
	}
	userManager := &UserManagerRepository{
		table:"user",
	}
	id ,err := userManager.Insert(user)
	if err != nil{
		panic(err)
	}
	user.ID = id
}
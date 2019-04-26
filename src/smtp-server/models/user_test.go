package models

import (
	"testing"
)

func TestCreateUser(t *testing.T){
	user := models.User{
		Email:"Wuduozhi@qq.com",
		Password:"wdz123",
		Nickname:"小智",
		Icon:"me.jpg",
		Create_at:122121,
	}

	_ = user.Create()

	user_login := models.Login("Wuduozhi@qq.com","wdz123")
}
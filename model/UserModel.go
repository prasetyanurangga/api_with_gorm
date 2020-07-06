package model

import(
	"time"
)

type (
	//User Is
	User struct{
		ID int `json:"id" gorm:"primary_key"`
		Nama string `json:"nama"`
		Gender string `json:"gender"`
		CreateAt time.Time `json:"create_at"`
	}

	Users []User
)
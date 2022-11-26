package models

import "time"

type Topping struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name      string               `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc      string               `json:"desc" gorm:"type:text" form:"desc"`
	Price     int                  `json:"price" form:"price" gorm:"type: int"`
	Image     string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty       int                  `json:"qty" form:"qty"`
	UserID    int                  `json:"user_id" form:"user_id"`
	User      UsersProfileResponse `json:"user"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type ToppingResponse struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name"`
	Desc   string               `json:"desc"`
	Price  int                  `json:"price"`
	Image  string               `json:"image"`
	Qty    int                  `json:"qty"`
	UserID int                  `json:"-"`
	User   UsersProfileResponse `json:"user"`
}

type ToppingUserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

func (ToppingResponse) TableName() string {
	return "toppings"
}

func (ToppingUserResponse) TableName() string {
	return "toppings"
}

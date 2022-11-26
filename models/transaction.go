package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Qty       int                  `json:"qty"`
	UsersID   int                  `json:"user_id"`
	Users     UsersProfileResponse `json:"userOrder" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status    string               `json:"status"`
	ProductID int                  `json:"product_id" gorm:"type: int"`
	Product   ProductResponse      `json:"order" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

package models

type Order struct {
	Id      uint    `json:"id"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Product Product ` json:"product" gorm:"foreignKey:ProductID"`

	Pid       uint `json:"pid"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"qty"`
}

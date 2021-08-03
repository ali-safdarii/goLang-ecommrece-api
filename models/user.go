package models

/*
type OwnModel struct {
	CreatedAt time.Time  `json:"-" `
	UpdatedAt time.Time  `json:"-" `
	DeletedAt *time.Time `json:"-"`
}*/

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	UserName string `json:"user_name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"` //`json:"-"` for first time maybe throw error so check on and change it to -> `json:"password"` then run again!
	RoleId   uint   `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleId"`
}

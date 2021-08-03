package models

type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float32 `json:"price"`
	IsExcite    bool    `json:"is_excite"`
}

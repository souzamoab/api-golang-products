package dto

type ProductCreateDTO struct {
	Title       string  `json:"user_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
}

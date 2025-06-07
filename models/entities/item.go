package entities

type Item struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    *string `json:"quantity"`
	Price       float64 `json:"price"`
	IsPerItem   bool    `json:"isPerItem"`
	UserId      *int    `json:"userId"`
}

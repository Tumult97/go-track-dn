package entities

type Item struct {
	BaseEntity
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    *string `json:"quantity"`
	Price       float64 `json:"price"`
	IsPerItem   bool    `json:"isPerItem"`
	LocationId  *int    `json:"locationId"`
}

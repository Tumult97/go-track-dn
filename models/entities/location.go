package entities

type Location struct {
	BaseEntity
	Name          string  `json:"name"`
	Description   *string `json:"description,omitempty"`
	AddressHome   *string `json:"addressHome,omitempty"`
	AddressStreet *string `json:"addressStreet,omitempty"`
	AddressSuburb *string `json:"addressSuburb,omitempty"`
	AddressCity   string  `json:"addressCity"`
}

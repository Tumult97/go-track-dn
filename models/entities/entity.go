package entities

type Entity interface {
	GetId() int
	SetId(id int)
	GetUserId() int
	SetUserId(id int)
}

type BaseEntity struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
}

func (b BaseEntity) GetId() int {
	return b.Id
}

func (b *BaseEntity) SetId(id int) {
	b.Id = id
}

func (b BaseEntity) GetUserId() int {
	return b.UserId
}

func (b *BaseEntity) SetUserId(id int) {
	b.UserId = id
}

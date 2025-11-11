package entities

import (
	"time"
)

type Entity interface {
	GetId() int
	SetId(id int)
	GetUserId() int
	SetUserId(id int)
	GetCreated() time.Time
	SetCreated(created time.Time)
	SetNowCreated()
}

type BaseEntity struct {
	Id         int       `json:"id"`
	UserId     int       `json:"userId"`
	Created    time.Time `json:"created"`
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

func (b *BaseEntity) GetCreated() time.Time {
	return b.Created
}

func (b *BaseEntity) SetCreated(created time.Time) {
	b.Created = created
}

func (b *BaseEntity) SetNowCreated() {
	b.Created = time.Now()
}

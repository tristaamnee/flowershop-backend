package model

import "time"

type User struct {
	ID       int64     `json:"id" bson:"id"`
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email" pg:",unique"`
	Password string    `json:"password" bson:"password"`
	Address  string    `json:"address" bson:"address"`
	Phone    string    `json:"phone" bson:"phone"`
	Birthday string    `json:"birthday" bson:"birthday"`
	Role     string    `json:"role" bson:"role"`
	Orders   []int64   `json:"orders" bson:"orders"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
}

const (
	ColUserID = "id"
	ColEmail  = "email"
)

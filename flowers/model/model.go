package model

import "time"

type Flower struct {
	ID         int64     `json:"id" bson:"id"`
	Name       string    `json:"name" bson:"name"`
	Stock      int64     `json:"stock" bson:"stock"`
	Price      float64   `json:"price" bson:"price"`
	Category   []string  `json:"category" bson:"category"`
	ExpiryDate time.Time `json:"expiry_date" bson:"expiry_date"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type Flowers struct {
	ID        int64     `json:"id" bson:"id"`
	FlowerIDs []int64   `json:"flower_ids" bson:"flower_ids"`
	UserID    int64     `json:"user_id" bson:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

const (
	ColFlowerID = "id"
	ColCategory = "category"
)

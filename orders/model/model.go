package model

type Order struct {
	ID           int64   `json:"id" bson:"id"`
	UserID       int64   `json:"user_id" bson:"user_id"`
	FlowerIDs    []int64 `json:"flower_ids" bson:"flower_ids"`
	PromotionIDs []int64 `json:"promotion_ids" bson:"promotion_ids"`
	Total        int64   `json:"total" bson:"total"`
	CreateAt     int64   `json:"create_at" bson:"create_at"`
	UpdateAt     int64   `json:"update_at" bson:"update_at"`
}

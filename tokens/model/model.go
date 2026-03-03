package model

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	ID           uuid.UUID        `pg:"type:uuid,pk" json:"id" bson:"_id"`
	CreationDate time.Time        `json:"creation_date" bson:"creation_date"`
	UserID       int64            `json:"user_id" bson:"user_id"`
	Activities   []*TokenActivity `json:"activities,omitempty" bson:"activities,omitempty"`
	_            struct{}         `pg:"_schema:tokens" json:"-" bson:"-"`
}

type TokenActivity struct {
	ID       int64     `json:"id" bson:"id"`
	TokenID  uuid.UUID `pg:"type:uuid,pk" json:"token_id" bson:"token_id"`
	Endpoint string    `json:"endpoint" bson:"endpoint"`
	UseAt    time.Time `json:"use_at" bson:"use_at"`
	_        struct{}  `pg:"_schema:token_activities" json:"-" bson:"-"`
}

const (
	ColTokenID = "id"
)

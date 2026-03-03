package repository

import (
	"github.com/go-pg/pg/v10"
	base "github.com/tristaamnee/flowershop-be/common/repository"
	"github.com/tristaamnee/flowershop-be/flowers/model"
)

func Create(db *pg.DB, flower *model.Flower) error {
	return base.Create(db, flower)
}

func GetByID(db *pg.DB, id int64) (*model.Flower, error) {
	return base.GetByCondition[model.Flower](db, model.ColFlowerID, id)
}

func GetByCategory(db *pg.DB, category string) (*model.Flower, error) {
	return base.GetByCondition[model.Flower](db, model.ColCategory, category)
}

func GetAllFlower(db *pg.DB) []*model.Flower {
	var flowers []*model.Flower
	base.GetWholeTable[model.Flower](db, &flowers)
	return flowers
}

func DeleteByID(db *pg.DB, id int64) error {
	return base.DeleteByCondition[model.Flower](db, model.ColFlowerID, id)
}

func UpdateFlower(db *pg.DB, id int64, data model.Flower) error {
	return base.Update(db, id, data)
}

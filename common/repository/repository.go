package repository

import (
	"log"

	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data interface{}) error {
	_, err := db.Model(data).Insert()
	if err != nil {
		log.Println("Error creating data: %v", err)
		return err
	}
	return nil
}

func GetWholeTable[T interface{}](db *pg.DB, Ts *[]*T, column ...string) {
	query := db.Model(Ts)

	if len(column) > 0 {
		query = query.Column(column...)
	}

	err := query.Select()
	if err != nil {
		log.Printf("Error selecting data: %v\n", err)
		return
	}
}

func GetByCondition[T interface{}](db *pg.DB, column string, value any) (*T, error) {
	data := new(T)
	err := db.Model(data).Where("? = ?", pg.Ident(column), value).Select()
	if err != nil {
		log.Println("Error getting data from %v ", column, " : %v", err)
		return nil, err
	}
	return data, nil
}

func Update(db *pg.DB, id int64, data interface{}) error {
	_, err := db.Model(data).Where("id = ?", id).Update()
	if err != nil {
		log.Println("Error updating data: %v", err)
		return err
	}
	return nil
}

func DeleteByCondition[T interface{}](db *pg.DB, column string, value any) error {
	var model T

	_, err := db.Model(&model).Where("? = ?", pg.Ident(column), value).Delete()
	if err != nil {
		log.Println("Error deleting data from %v ", column, " : %v", err)
		return err
	}
	return nil
}

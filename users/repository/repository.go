package repository

import (
	"github.com/go-pg/pg/v10"
	base "github.com/tristaamnee/flowershop-be/common/repository"
	"github.com/tristaamnee/flowershop-be/users/model"
)

func Create(db *pg.DB, user *model.User) error {
	return base.Create(db, user)
}

func GetByID(db *pg.DB, id int64) (*model.User, error) {
	return base.GetByCondition[model.User](db, model.ColUserID, id)
}

func GetByEmail(db *pg.DB, email string) (*model.User, error) {
	return base.GetByCondition[model.User](db, model.ColEmail, email)
}

func DeleteByEmail(db *pg.DB, email string) error {
	return base.DeleteByCondition[model.User](db, model.ColEmail, email)
}

func Update(db *pg.DB, id int64, data model.User) error {
	return base.Update(db, id, data)
}

func DeleteByID(db *pg.DB, id int64) error {
	return base.DeleteByCondition[model.User](db, model.ColUserID, id)
}

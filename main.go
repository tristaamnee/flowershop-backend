package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/common/db"
	"github.com/tristaamnee/flowershop-be/common/utils/configuration"
	flowersModel "github.com/tristaamnee/flowershop-be/flowers/model"
	ordersModel "github.com/tristaamnee/flowershop-be/orders/model"
	usersModel "github.com/tristaamnee/flowershop-be/users/model"
)

func main() {

	r := gin.Default()

	conf, err := configuration.Get("db")
	if err != nil {
		log.Fatalln("error reading configuration file:", err)
	}

	database := db.ConnectDatabase(conf.(configuration.Db))
	defer database.Close()

	if err := r.Run(":8080"); err != nil {
		log.Fatalln("error running server:", err)
	}

}

func createDBTables(err error, database *pg.DB) {

	//Create the 'users' table
	err = db.CreateTable(database, (*usersModel.User)(nil))
	if err != nil {
		log.Fatalf("Error creating users table:", err)
	}

	//Create the 'flower' table
	err = db.CreateTable(database, (*flowersModel.User)(nil))
	if err != nil {
		log.Fatalf("Error creating flowers table:", err)
	}

	//Create the 'order' table
	err = db.CreateTable(database, (*ordersModel.Order)(nil))
	if err != nil {
		log.Fatalf("Error creating orders table:", err)
	}
}

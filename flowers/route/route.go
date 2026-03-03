package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/flowers/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.POST("/createflower", handler.CreateFlower(db))
	r.GET("/flowers/:id", handler.GetFlowerByID(db))
	r.GET("/flowers", handler.GetAllFlowerHandler(db))
}

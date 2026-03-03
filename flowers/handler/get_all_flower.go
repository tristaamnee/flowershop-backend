package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/flowers/repository"
)

func GetAllFlowerHandler(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		flowerData := repository.GetAllFlower(db)
		c.JSON(http.StatusOK, flowerData)
	}
}

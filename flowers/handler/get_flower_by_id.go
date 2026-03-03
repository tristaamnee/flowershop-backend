package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/flowers/model"
	"github.com/tristaamnee/flowershop-be/flowers/repository"
)

func GetFlowerByID(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param(model.ColFlowerID)
		flowerID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid flower id",
			})
			return
		}

		flowerData, err := repository.GetByID(db, flowerID)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "flower not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"flower": flowerData,
		})
	}
}

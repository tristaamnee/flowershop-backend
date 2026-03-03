package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/flowers/model"
	"github.com/tristaamnee/flowershop-be/flowers/repository"
)

func DeleteFlowerById(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param(model.ColFlowerID)
		flowerID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid flower id",
			})
			return
		}

		err = repository.DeleteByID(db, flowerID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Flower not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Flower deleted",
		})
	}
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/flowers/model"
	"github.com/tristaamnee/flowershop-be/flowers/repository"
)

func GetFlowerByCategory(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryParam := c.Param(model.ColCategory)
		flowerData, err := repository.GetByCategory(db, categoryParam)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"flowers": flowerData,
		})
	}
}

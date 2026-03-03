package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/flowers/model"
	"github.com/tristaamnee/flowershop-be/flowers/repository"
)

func CreateFlower(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input *model.Flower
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input",
			})
			return
		}

		err := repository.Create(db, input)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create flower", "message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Successfully created flower",
			"flower":  input,
		})
	}
}

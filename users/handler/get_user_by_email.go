package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/users/model"
	"github.com/tristaamnee/flowershop-be/users/repository"
)

func GetUserByEmail(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param(model.ColEmail)
		userData, err := repository.GetByEmail(db, email)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "User not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": userData,
		})
	}
}

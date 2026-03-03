package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/users/model"
	"github.com/tristaamnee/flowershop-be/users/repository"
)

func DeleteUserById(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param(model.ColUserID)
		userID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid user id",
			})
			return
		}

		//delete user using id
		err = repository.DeleteByID(db, userID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "User not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "User deleted",
		})
	}
}

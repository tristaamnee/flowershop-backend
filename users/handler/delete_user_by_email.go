package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/users/model"
	"github.com/tristaamnee/flowershop-be/users/repository"
)

func DeleteUserByEmail(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param(model.ColEmail)

		//delete user by email
		err := repository.DeleteByEmail(db, email)
		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted"})
	}
}

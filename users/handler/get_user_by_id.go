package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/users/model"
	"github.com/tristaamnee/flowershop-be/users/repository"
)

func GetUserByID(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param(model.ColUserID)
		userID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"Error": "invalid user id"})
			return
		}

		// get user using ID
		userData, err := repository.GetByID(db, userID)
		if err != nil {
			c.JSON(400, gin.H{"Error": "user not found"})
			return
		}
		c.JSON(200, gin.H{"user": userData})
	}
}

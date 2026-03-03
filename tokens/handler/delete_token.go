package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/tristaamnee/flowershop-be/tokens/repository"
)

func DeleteToken(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenIDParameter := c.Param("tokenID")
		tokenID, err := uuid.Parse(tokenIDParameter)
		if err != nil {
			fmt.Println("Error parsing UUID:", err)
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}

		err = repository.DeleteTokenActivity(db, tokenID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Token activity doesn't exist"})
			return
		}

		err = repository.DeleteToken(db, tokenID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Token not found"})
			return
		}
		c.JSON(200, gin.H{"result": "token has bene deleted"})

	}
}

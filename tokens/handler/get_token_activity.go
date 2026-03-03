package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/tristaamnee/flowershop-be/tokens/repository"
)

func GetTokenActivity(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenIDParam := c.Param("tokenID")

		parsedUUID, err := uuid.Parse(tokenIDParam)
		if err != nil {
			fmt.Println("Error parsing UUID:", err)
			c.JSON(400, gin.H{"error": err})
			c.Abort()
			return
		}

		tokenActivity, err := repository.GetTokenActivity(db, parsedUUID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Token doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"token_activities": tokenActivity})
	}
}

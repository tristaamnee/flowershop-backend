package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/common/security/jwt"
	"github.com/tristaamnee/flowershop-be/users/repository"
)

func LoginUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.MustGet("username").(string)
		user, err := repository.GetByEmail(db, username)
		if err != nil {
			c.JSON(401, gin.H{"message": err})
			return
		}
		token, err := jwt.GenerateToken(db, user)
		if err != nil {
			c.JSON(401, gin.H{"error": err})
			return
		}

		c.JSON(200, gin.H{
			"token": token,
		})
	}
}

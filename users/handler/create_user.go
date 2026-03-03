package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/common/repository"
	"github.com/tristaamnee/flowershop-be/users/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input *model.User
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		//password hash
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}
		input.Password = string(hashedPassword)

		//try to add new user to db
		err = repository.Create(db, input)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user", "message": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User created successfully", "user": input})

	}
}

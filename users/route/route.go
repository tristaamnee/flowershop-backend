package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/tristaamnee/flowershop-be/users/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.POST("/register", handler.CreateUser(db))
}

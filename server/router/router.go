package router

import (
	"bot/api"
	"bot/auth"
	"bot/util"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB, userlist []util.Userlist) {
	router := gin.Default()
	router.POST("/bot/newuser", api.POST_NewUser(db))
	router.GET("/bot/user", auth.AuthMiddleware(userlist), api.HandleQuestionRequest)
	router.Run("localhost:8091")
}

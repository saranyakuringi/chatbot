package auth

import (
	"bot/model"
	"bot/util"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userlist []util.Userlist

func UsersList(db *sql.DB) ([]util.Userlist, error) {
	usersdata, err := model.UserList(db)
	if err != nil {
		fmt.Println("Error in userlist function", err)
		return nil, err
	}
	userlist = usersdata
	for _, value := range usersdata {
		fmt.Printf("User:%s,Password:%s\n", value.Username, value.Password)
	}
	fmt.Println("user in userslist output", userlist)
	return userlist, err
}

func AuthMiddleware(userlist []util.Userlist) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		fmt.Printf("User:%s,Password:%s\n", username, password)

		if ok {
			for _, user := range userlist {
				if username == user.Username && password == user.Password {
					c.Next()
					fmt.Println("Authentication successful.")
					return
				}
			}
		}
		log.Println("Authentication failed.")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unauthorised"})
		c.Abort()
	}
}

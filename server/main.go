package main

import (
	"bot/auth"
	"bot/controller"
	"bot/router"
	"database/sql"
	"fmt"
	"sync"
)

var db *sql.DB
var dbMutex sync.Mutex

func initDB() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	// Check if the database is already initialized
	if db != nil {
		return
	}

	var err error
	db, err = controller.Connecting_database()
	if err != nil {
		fmt.Println("Error in connecting to database", err)
		panic("Failed to connect to the database")
	}
}
func main() {
	fmt.Println("Initializing the bot....")

	// Initialize database connection
	initDB()

	userlist, err := auth.UsersList(db)
	if err != nil {
		fmt.Println("Error in getting user list", err)
		panic("Failed to get user list")
	}
	//start the router
	router.Router(db, userlist)

}

package model

import (
	"bot/util"
	"database/sql"
	"fmt"
	"log"
)

// insert new user
func Insert_User(db *sql.DB, user util.Userlist) (util.Userlist, error) {
	query := `
	Insert into users
	(username,password) values ($1,$2)
	returning username,password
	`
	err := db.QueryRow(query, user.Username, user.Password).Scan(&user.Username, &user.Password)
	if err != nil {
		log.Println("Error in User Insert query", err)
		return util.Userlist{}, err
	}
	fmt.Println(user)
	return user, err
}

// fetch all the existing users
func UserList(db *sql.DB) ([]util.Userlist, error) {
	rows, err := db.Query(`Select * from users`)
	if err != nil {
		log.Println("Error in users search query", err)
		return nil, err
	}

	//iterating through rows
	var user util.Userlist
	var output []util.Userlist
	for rows.Next() {
		err = rows.Scan(
			&user.Username,
			&user.Password,
		)
		output = append(output, user)
	}
	fmt.Println("users list", output)
	return output, err
}

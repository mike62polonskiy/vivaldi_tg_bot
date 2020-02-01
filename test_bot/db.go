package main

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var dbname = os.Getenv("DB_NAME")
var sslmode = os.Getenv("SSLMODE")

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

func userReg(chatID int64, username string) error {

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating SQL command
	data := `INSERT INTO tg_bot_users(chat_id, username) VALUES($1, $2);`

	//Execute SQL command in database
	if _, err = db.Exec(data, chatID, `@`+username); err != nil {
		return err
	}

	return nil
}

func checkExistUser(username string) string {

	var user string

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("ERROR: failed connect to db")
	}
	defer db.Close()

	//Запрос на выборку существующего пользователя.
	sqlState := `SELECT username FROM tg_bot_users where username = $1`
	
	row := db.QueryRow(sqlState, username)

	switch err := row.Scan(&user); err {
	case sql.ErrNoRows:
	  return ""
	case nil:
	  fmt.Println("User exist")
	default:
	  panic(err)
	}

	return user
	
}

func getPlaces() []string {

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("ERROR: failed to connect to db")
	}
	defer db.Close()

	sqlState := `SELECT group_tag FROM vk_data_grub_vkgroups`
	
	rows, err := db.Query(sqlState)

	if err != nil {
		fmt.Println(err)
	}

	grTags :=[]string{}
	for rows.Next() {
		var grTag string
		rows.Scan(&grTag)
		grTags = append(grTags, grTag)
	}

	return grTags
}
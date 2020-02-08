package main

import (
	"os"
	"fmt"
	"reflect"
	"encoding/json"
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

func checkExist(sqlData string, varName string) string {

	var param string

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("ERROR: failed connect to db")
	}
	defer db.Close()

	row := db.QueryRow(sqlState, varName)

	switch err := row.Scan(&param); err {
	case sql.ErrNoRows:
	  return ""
	case nil:
	  fmt.Println("Row exist")
	default:
	  panic(err)
	}

	return param
}

func sqlToJSON(sqlState string) ([]byte, error) {

	var objects []map[string]interface{}

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("ERROR: failed to connect to db")
	}
	defer db.Close()
	
	rows, err := db.Query(sqlState)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		columns, err := rows.ColumnTypes()
		if err != nil {
			fmt.Println(err)
		}

		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			object[column.Name()] = reflect.New(column.ScanType()).Interface()
			values[i] = object[column.Name()]
		}
	
		err = rows.Scan(values...)
		if err != nil {
			fmt.Println(err)
		}

		objects = append(objects, object)
	}

	return json.MarshalIndent(objects, "", "\t")

}

func updateUserCity(city string, username string) error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	data := `UPDATE tg_bot_users SET city_id = (SELECT id FROM tg_bot_cities WHERE city ILIKE $1) WHERE username = $2;`
	
	if _, err = db.Exec(data, city, `@`+username); err != nil {
		return err
	}

	return nil

}

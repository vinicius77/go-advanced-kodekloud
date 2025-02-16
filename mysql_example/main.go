package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DBName = "learning"
const DBUser = "root"
const DBPASS = "root"

type Data struct {
	id   int
	name string
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}

}

func main() {
	connectionStr := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPASS, DBName)
	db, err := sql.Open("mysql", connectionStr)
	checkError(err)

	defer db.Close()

	result, err := db.Exec("INSERT into data values(4, 'simple insert')")
	checkError(err)

	lastInsertedId, err := result.LastInsertId()
	checkError(err)
	fmt.Println("last inserted id", lastInsertedId)

	rowsAffected, err := result.RowsAffected()
	checkError(err)
	fmt.Println("row affected", rowsAffected)

	rows, err := db.Query("SELECT * from data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}
}

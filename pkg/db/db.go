package db

import (
	"database/sql"
	"fmt"

	 _ "github.com/mattn/go-sqlite3"
)

func Init() {
	db, err := sql.Open("sqlite3", "D:\\Swapnil D\\Projects\\ebook-creator-backend\\pkg\\db\\db.sqlite")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error one tripped")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Two tripped")
		return
	}
	fmt.Println("Ping")
}

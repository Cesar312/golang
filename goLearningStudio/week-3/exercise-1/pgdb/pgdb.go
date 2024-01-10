package pgdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	connectionString := "user=postgres dbname=postgres password=Ljie@H6PNE.uF!y2wm host=localhost sslmode=disable"
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func PostUpdate() {
	db := GetDB()
	_, err := db.Exec(`INSERT INTO studiodb(name, title) VALUES('Cesar Martinez', 'Data Scientist')`)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

package pgdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBRecord struct {
	Studio_Id  string
	Name       string
	Title      string
	Created_At string
}

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

func GetRecords() []DBRecord {
	var DBRecords []DBRecord
	db := GetDB()

	rows, err := db.Query(`SELECT studio_id, name, title, created_at FROM studiodb;`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var dbreturn DBRecord

		if err := rows.Scan(&dbreturn.Studio_Id, &dbreturn.Name, &dbreturn.Title, &dbreturn.Created_At); err != nil {
			fmt.Println(err)
		}
		DBRecords = append(DBRecords, dbreturn)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	return DBRecords

}

package main

import (
	"fmt"
	pgdb "wk3e3/pgdb"
)

func main() {
	DBRecord := pgdb.GetRecords()
	fmt.Println(DBRecord)

	fmt.Println("Database Retrieval Completed")
}

package main

import (
	"fmt"
	pgdb "week3/pgdb"
)

func main() {
	pgdb.PostUpdate()
	fmt.Println("Database update is complete.")
}

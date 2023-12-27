package main

import (
	"fmt"
	pgdb "wk3e2/pgdb"
)

func main() {

	// pgdb.PostUpdate()
	pgdb.PostUpdate2("John Smith", "Software Engineer")

	fmt.Println("Database update is complete.")
}

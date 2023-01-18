package main

import (
	"fmt"
	"os"

	"cool-api/database"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		registerRoute(db)
	}
}

package main

import (
	"fmt"
	"os"

	d "cool-api/database"
)

func main() {
	db, err := d.Connect()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		registerRoute(db)
	}
}

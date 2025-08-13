package main

import (
	"log"
	"server/db"
)

func main() {
	_, err := db.Newdatabase()
	if err != nil {
		log.Fatalf("could not initialize DB conn. %s", err)
	}
}

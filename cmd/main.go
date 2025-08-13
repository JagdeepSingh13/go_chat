package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.Newdatabase()
	if err != nil {
		log.Fatalf("could not initialize DB conn. %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSrv := user.NewService(userRep)
	userHandler := user.NewHandler(userSrv)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}

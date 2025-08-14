package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
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

	hub := ws.NewHub()
	go hub.Run()
	wsHandler := ws.NewHandler(hub)

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}

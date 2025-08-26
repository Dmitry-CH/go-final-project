package main

import (
	"log"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
	"github.com/Dmitry-CH/go-final-project/pkg/server"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalf("Start db error: %s", err.Error())
		return
	}
	defer db.Close()

	err = server.Run()
	if err != nil {
		log.Fatalf("Start server error: %s", err.Error())
	}
}

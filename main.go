package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
	"github.com/joho/godotenv"
)

const webDir = "web"

var port = 7540

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	err = db.Init()
	if err != nil {
		log.Fatalf("Start db error: %s", err.Error())
		return
	}
	defer db.Close()

	envPort := os.Getenv("TODO_PORT")
	if len(envPort) > 0 {
		if eport, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(eport)
		}
	}

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Start server error: %s", err.Error())
	}
}

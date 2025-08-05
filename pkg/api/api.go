package api

import "net/http"

const webDir = "web"

func Init() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.HandleFunc("GET /api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
}

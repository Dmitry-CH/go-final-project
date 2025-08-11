package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const webDir = "web"

func writeErrJson(w http.ResponseWriter, msg string) {
	resp := []byte(fmt.Sprintf(`{"error": "%s"}`, msg))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(resp)
}

func writeJson(w http.ResponseWriter, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		writeErrJson(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(resp)
}

func Init() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.HandleFunc("GET /api/nextdate", nextDateHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task", taskHandler)
}

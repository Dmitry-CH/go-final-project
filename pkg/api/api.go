package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const webDir = "web"

type ErrResp struct {
	Error string `json:"error"`
}

func Init() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.HandleFunc("GET  /api/nextdate", nextDateHandler)
	http.HandleFunc("     /api/task", taskHandler)
	http.HandleFunc("POST /api/task/done", taskDoneHandler)
	http.HandleFunc("GET  /api/tasks", tasksHandler)
}

func writeJson(w http.ResponseWriter, data any, statusCode ...int) {
	resp, err := json.Marshal(data)
	if err != nil {
		resp = []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	}
	w.Write(resp)
}

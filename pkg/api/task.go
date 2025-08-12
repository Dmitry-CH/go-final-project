package api

import (
	"errors"
	"net/http"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

var ErrGetTask = errors.New("ошибка задача не найдена")

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	rId := r.URL.Query().Get("id")

	task, err := db.GetTask(rId)
	if err != nil {
		writeJson(w, ErrResp{ErrGetTask.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, task)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	}
}

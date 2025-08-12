package api

import (
	"net/http"
	"time"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	rSearch := r.URL.Query().Get("search")

	date := ""
	search := rSearch

	t, err := time.Parse("02.01.2006", search)
	if err == nil {
		date = t.Format(DATE_FORMAT)
		search = ""
	}

	tasks, err := db.GetTasks(50, search, date)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, TasksResp{tasks})
}

package api

import (
	"net/http"
	"time"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	rId := r.URL.Query().Get("id")

	task, err := db.GetTask(rId)
	if err != nil {
		writeJson(w, ErrResp{ErrGetTask.Error()}, http.StatusInternalServerError)
		return
	}

	if len(task.Repeat) == 0 {
		err = db.DeleteTask(rId)
		if err != nil {
			writeJson(w, ErrResp{ErrGetTask.Error()}, http.StatusInternalServerError)
			return
		}
	} else {
		now := time.Now()

		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJson(w, ErrResp{ErrGetTask.Error()}, http.StatusInternalServerError)
			return
		}

		err = db.UpdateTaskDate(next, rId)
		if err != nil {
			writeJson(w, ErrResp{ErrGetTask.Error()}, http.StatusInternalServerError)
			return
		}
	}

	writeJson(w, EmpResp{})
}

package api

import (
	"net/http"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	rId := r.URL.Query().Get("id")

	task, err := db.GetTask(rId)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, task)
}

package api

import (
	"net/http"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	rId := r.URL.Query().Get("id")

	err := db.DeleteTask(rId)
	if err != nil {
		writeJson(w, ErrResp{ErrGetTask.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, EmpResp{})
}

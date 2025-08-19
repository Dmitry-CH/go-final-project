package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task db.Task

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		writeJson(w, ErrResp{"ошибка десериализации JSON"}, http.StatusBadRequest)
		return
	}
	if len(task.Title) == 0 {
		writeJson(w, ErrResp{"ошибка не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, EmpResp{})
}

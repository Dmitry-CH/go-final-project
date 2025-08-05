package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

type successResponse struct {
	ID int64 `json:"id"`
}

type failedResponse struct {
	Error string `json:"error"`
}

func checkDate(task *db.Task) error {
	now := time.Now()
	now = now.Truncate(24 * time.Hour) // сбрасываем время до 00:00:00 +0000 для коректной работы afterNow(now, t).

	if len(task.Date) == 0 {
		task.Date = now.Format(DATE_FORMAT)
	}

	t, err := time.Parse(DATE_FORMAT, task.Date)
	if err != nil {
		return errors.New("ошибка дата представлена в формате, отличном от 20060102")
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format(DATE_FORMAT)
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return errors.New("ошибка правило повторения указано в неправильном формате")
			}

			// в противном случае, берём вычисленную ранее следующую дату
			task.Date = next
		}
	}

	return nil
}

func writeJson(w http.ResponseWriter, data any, code int) {
	resp, err := json.Marshal(data)
	if err != nil {
		resp = []byte(`{"error": "ошибка сериализации JSON"}`)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(resp)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task db.Task

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, failedResponse{err.Error()}, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		writeJson(w, failedResponse{"ошибка десериализации JSON"}, http.StatusBadRequest)
		return
	}
	if len(task.Title) == 0 {
		writeJson(w, failedResponse{"ошибка не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, failedResponse{err.Error()}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, failedResponse{err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, successResponse{id}, http.StatusOK)
}

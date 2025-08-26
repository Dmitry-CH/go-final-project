package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task db.Task

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}
	if len(task.Title) == 0 {
		writeJson(w, ErrResp{"title required"}, http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, IDResp{id})
}

func checkDate(task *db.Task) error {
	now := time.Now()
	now = now.Truncate(24 * time.Hour) // сбрасываем время до 00:00:00 +0000 для коректной работы afterNow(now, t).

	if len(task.Date) == 0 {
		task.Date = now.Format(DATE_FORMAT)
	}

	t, err := time.Parse(DATE_FORMAT, task.Date)
	if err != nil {
		return errors.New("date is presented in a format other than 20060102")
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format(DATE_FORMAT)
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return errors.New("repetition rule is specified in the wrong format")
			}

			// в противном случае, берём вычисленную ранее следующую дату
			task.Date = next
		}
	}

	return nil
}

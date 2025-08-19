package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", errors.New("ошибка в параметре 'repeat' — пустая строка")
	}

	date, err := time.Parse(DATE_FORMAT, dstart)
	if err != nil {
		return "", errors.New("ошибка время в параметре 'dstart' не может быть преобразовано в корректную дату")
	}

	rule := strings.Split(repeat, " ")
	var interval []int

	switch rule[0] {
	case "d":
		if len(rule) < 2 {
			return "", errors.New("ошибка указан неверный формат 'repeat' - не указан интервал в днях")
		}

		days, err := strconv.Atoi(rule[1])
		if err != nil {
			return "", errors.New("ошибка указан неверный формат 'repeat' - недопустимый символ")
		}
		if days > 400 {
			return "", errors.New("ошибка указан неверный формат 'repeat' - превышен максимально допустимый интервал")
		}

		interval = []int{0, 0, days}
	case "y":
		interval = []int{1, 0, 0}
	default:
		return "", errors.New("ошибка указан неверный формат 'repeat' - неподдерживаемый формат")
	}

	for {
		date = addDate(date, interval...)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format(DATE_FORMAT), nil
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	rNow := r.FormValue("now")
	rDate := r.FormValue("date")
	rRepeat := r.FormValue("repeat")

	now := time.Now()
	if len(rNow) > 0 {
		pNow, err := time.Parse(DATE_FORMAT, rNow)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		now = pNow
	}

	nextDate, err := NextDate(now, rDate, rRepeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(nextDate))
	if err != nil {
		http.Error(w, "ошибка сервера", http.StatusInternalServerError)
		return
	}
}

func addDate(date time.Time, s ...int) time.Time {
	return date.AddDate(s[0], s[1], s[2])
}

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

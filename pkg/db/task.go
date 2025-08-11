package db

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date, title, comment, repeat)
					VALUES (:date, :title, :comment, :repeat);`

	res, err := db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return id, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, nil
}

func Tasks(limit int, search, date string) ([]*Task, error) {
	tasks := make([]*Task, 0)
	query := `SELECT * FROM scheduler ORDER BY date LIMIT :limit;`

	if len(search) > 0 {
		query = `SELECT * FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit;`
		search = fmt.Sprintf(`%%%s%%`, search)
	}

	if len(date) > 0 {
		query = `SELECT * FROM scheduler WHERE date = :date LIMIT :limit`
	}

	rows, err := db.Query(query,
		sql.Named("limit", limit),
		sql.Named("search", search),
		sql.Named("date", date),
	)
	if err != nil {
		return []*Task{}, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return []*Task{}, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

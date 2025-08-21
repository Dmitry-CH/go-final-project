package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Dmitry-CH/go-final-project/pkg/db"
	"github.com/golang-jwt/jwt"
)

const webDir = "web"

type EmpResp struct{}

type ErrResp struct {
	Error string `json:"error"`
}

type IDResp struct {
	ID int64 `json:"id"`
}

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

type TokenResp struct {
	Token string `json:"token"`
}

func Init() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.HandleFunc("GET  /api/nextdate", nextDateHandler)
	http.HandleFunc("POST /api/signin", signinHandler)
	http.HandleFunc("     /api/task", auth(taskHandler))
	http.HandleFunc("POST /api/task/done", auth(taskDoneHandler))
	http.HandleFunc("GET  /api/tasks", auth(tasksHandler))
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ePass := os.Getenv("TODO_PASSWORD")
		if len(ePass) > 0 {
			var rToken string
			var valid bool

			cookie, err := r.Cookie("token")
			if err == nil {
				rToken = cookie.Value
			}

			jwtToken, err := jwt.Parse(rToken, func(t *jwt.Token) (any, error) {
				return secretKey, nil
			})
			if err != nil {
				writeJson(w, ErrResp{"failed to parse token"}, http.StatusUnauthorized)
				return
			}

			if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
				pass, ok := claims["pass"]
				if ok {
					valid = pass == generateSum(ePass)
				}
			}

			if !valid {
				writeJson(w, ErrResp{"authentification required"}, http.StatusUnauthorized)
				return
			}
		}

		next(w, r)
	})
}

func writeJson(w http.ResponseWriter, data any, statusCode ...int) {
	resp, err := json.Marshal(data)
	if err != nil {
		resp = []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	}
	w.Write(resp)
}

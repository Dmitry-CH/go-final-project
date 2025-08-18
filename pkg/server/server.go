package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Dmitry-CH/go-final-project/pkg/api"
)

var port = 7540

func Run() error {
	ePort := os.Getenv("TODO_PORT")
	if len(ePort) > 0 {
		if eport, err := strconv.ParseInt(ePort, 10, 32); err == nil {
			port = int(eport)
		}
	}

	api.Init()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

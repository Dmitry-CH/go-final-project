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
	envPort := os.Getenv("TODO_PORT")
	if len(envPort) > 0 {
		if eport, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(eport)
		}
	}

	api.Init()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dmitry-CH/go-final-project/pkg/api"
	"github.com/Dmitry-CH/go-final-project/pkg/config"
)

var port = 7540

func Run() error {
	conf := config.New()

	ePort := conf.Port
	if len(ePort) > 0 {
		if eport, err := strconv.ParseInt(ePort, 10, 32); err == nil {
			port = int(eport)
		}
	}

	api.Init()

	fmt.Printf("Server is running on port %d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello API, current Time:%s", time.Now().UTC().Format(time.RFC3339))
	})
	return router
}

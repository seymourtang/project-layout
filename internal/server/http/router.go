package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		clientCountry := request.Header.Get("Cf-Ipcountry")
		clientIP := request.Header.Get("Cf-Connecting-Ip")
		log.Printf("request,clientIP:%s,clientCountry:%s", clientIP, clientCountry)
		body := struct {
			Code    int                    `json:"code"`
			Message string                 `json:"message"`
			Data    map[string]interface{} `json:"data"`
			Time    string                 `json:"time"`
		}{
			Code:    0,
			Message: "Success",
			Data: map[string]interface{}{
				"userAgent":     request.UserAgent(),
				"clientCountry": clientCountry,
				"clientIP":      clientIP,
			},
			Time: time.Now().UTC().Format(time.RFC3339Nano),
		}
		json.NewEncoder(writer).Encode(body)
	})
	return router
}

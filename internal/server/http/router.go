package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
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
	router.GET("/chat", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		conn, _, _, err := ws.UpgradeHTTP(request, writer)
		if err != nil {
			log.Printf("unable to upgrade websocket,err:%s", err.Error())
			fmt.Fprintf(writer, "unable to upgrade websocket,err:%s", err.Error())
			return
		}
		clientCountry := request.Header.Get("Cf-Ipcountry")
		clientIP := request.Header.Get("Cf-Connecting-Ip")
		log.Printf("new connection,IP:%s,Country:%s", clientIP, clientCountry)
		go func() {
			defer conn.Close()
			defer func() {
				log.Printf("connection closed,IP:%s,Country:%s", clientIP, clientCountry)
			}()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Printf("failed to read data,conn:%s,err:%s", conn.RemoteAddr().String(), err.Error())
					return
				}
				log.Printf("received data from [%s],msg:%s", conn.RemoteAddr(), msg)
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					log.Printf("failed to write data,conn:%s,err:%s", conn.RemoteAddr().String(), err.Error())
					return
				}
			}
		}()
	})
	return router
}

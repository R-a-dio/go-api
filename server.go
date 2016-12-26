package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func startServer() error {
	srv := http.Server{
		Addr:    "localhost:5500",
		Handler: http.HandlerFunc(APIHandler),
	}

	return srv.ListenAndServe()
}

func APIHandler(rw http.ResponseWriter, r *http.Request) {
	var js jsonMain
	js.Main = apiMain.Load().(API)
	js.Main.CurrentTime = time.Now().Unix()
	js.Main.Queue = apiQueue.Load().([]ListEntryAPI)
	js.Main.LastPlayed = apiLastPlayed.Load().([]ListEntryAPI)

	e := json.NewEncoder(rw)
	e.SetEscapeHTML(false)
	err := e.Encode(js)
	if err != nil {
		log.Printf("json encoding error: %s", err)
	}
}

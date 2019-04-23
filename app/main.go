package main

import (
	"encoding/json"
	"google.golang.org/appengine"
	"net/http"
)


type Ping struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var p Ping
	p.Message = "Hello World!!"
	res, err := json.Marshal(p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	var ping Ping
	ping.Message = "Hello World!!"

	http.HandleFunc("/", Handler)

	// for GAE/Go
	appengine.Main()

	/*
	// for localhost
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
	*/
}

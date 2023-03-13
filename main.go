package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/valve1", valve1)
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func index(w http.ResponseWriter, c *http.Request) {
	if c.Method == "GET" {
		log.Println("index invoked")
		http.Error(w, "Not Found", 404)
	}
}

type valve struct {
	Val string `json:"mili"`
}

func valve1(w http.ResponseWriter, c *http.Request) {
	if c.Method == "POST" {
		data := json.NewDecoder(c.Body)
		var t valve
		err := data.Decode(&t)
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid body", 404)
			return
		}
		log.Println(t.Val)
	}
}

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"mcu-server/config"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var conf *config.Config

func main() {
	data, err := config.Load()
	config.LogErr(err)
	conf = &data
	config.LogErr(connect())
	http.HandleFunc("/", index)
	http.HandleFunc("/flow", valve1)
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(conf.Host, nil))
}

func connect() error {
	db, err := sql.Open("postgres", conf.DbUrl)
	if err != nil {
		return err
	}
	if err2 := db.Ping(); err2 != nil {
		return err2
	}
	return nil
}

func index(w http.ResponseWriter, c *http.Request) {
	if c.Method == "GET" {
		log.Println("index invoked")
		http.Error(w, "Not Found", 404)
	}
}

type valve struct {
	Val    string `json:"value"`
	Device string `json:"device"`
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
		db.Exec("insert into flow_one (mili) values ($1)", t.Val)
		log.Println(t)
	}
}

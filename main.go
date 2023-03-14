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
	if conf.Use {
		config.LogErr(connect())
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/flow", flow)
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

func flow(w http.ResponseWriter, c *http.Request) {
	if c.Method == "POST" {
		data := json.NewDecoder(c.Body)
		var t config.Sensor
		err := data.Decode(&t)
		config.LogIgnore(err)
		log.Println(t)
		if conf.Use {
			if t.Device == "flow1" {
				db.Exec("insert into flow_one (mili) values ($1)", t.Val)
			} else {
				db.Exec("insert into flow_two (mili) values ($1)", t.Val)
			}
		}
	}
}

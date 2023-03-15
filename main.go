package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"mcu-server/config"
	"net/http"
  "strconv"

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
	conn, err := sql.Open("postgres", conf.DbUrl)
  db = conn
	if err != nil {
		return err
	}
	if err2 := db.Ping(); err2 != nil {
		return err2
	}
  log.Println("database pool established")
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
    float, err := strconv.ParseFloat(t.Val, 32)
    config.LogErr(err)
		if conf.Use {
			if t.Device == "flow1" {
        _,err := db.Exec("insert into flow_sens (sens_id,flow) values (1,$1)", float)
        config.LogErr(err)
			} else {
        _,err := db.Exec("insert into flow_sens (sens_id,flow) values (2,$1)", float)
        config.LogErr(err)
			}
		}
		w.WriteHeader(http.StatusOK)
	}
}

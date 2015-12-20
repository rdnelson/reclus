package main

import (
	"fmt"
	"net/http"

	"github.com/rdnelson/reclus/backends"
	"github.com/rdnelson/reclus/config"
	"github.com/rdnelson/reclus/log"
)

func main() {
	log.Log.Print("Starting Reclus Issue Tracker...")

	if err := config.Load(); err != nil {
		log.Log.Fatal(err)
	}

	log.Log.Level = config.Cfg.Logging.GetLogLevel()

	db, err := backends.NewDatabase()

	if err != nil {
		log.Log.Fatal(err)
	}

	defer db.Close()

	if err = setupAuth(db); err != nil {
		log.Log.Fatal(err)
	}

	mux := NewRouter()

	log.Log.Print("Starting HTTP server")

	log.Log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Cfg.Server.Hostname, config.Cfg.Server.Port), mux))
}

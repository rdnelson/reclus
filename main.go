package main

import (
	"fmt"
	"net/http"

	"github.com/rdnelson/reclus/backends"
	"github.com/rdnelson/reclus/config"
	"github.com/rdnelson/reclus/datamodel"
	"github.com/rdnelson/reclus/log"

	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/register"

	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
)

var (
	authManager = authboss.New()
)

func setupAuth(db backends.Database) error {

	if err := InitializeCookieStore(); err != nil {
		return err
	}

	if err := InitializeSessionStore(); err != nil {
		return err
	}

	authManager.Storer = NewUserRepo(db)
	authManager.SessionStoreMaker = NewSessionStore
	authManager.CookieStoreMaker = NewCookieStore
	authManager.MountPath = "/auth"
	authManager.RootURL = "http://localhost:8080"
	authManager.LogWriter = log.LogWriter{log.Log}

	authManager.XSRFName = "csrf_token"
	authManager.XSRFMaker = func(_ http.ResponseWriter, req *http.Request) string {
		return nosurf.Token(req)
	}

	return authManager.Init()
}

func main() {
	log.Log.Level = log.DebugLevel
	log.Log.Print("Starting Reclus Issue Tracker...")

	if err := config.Load(); err != nil {
		log.Log.Fatal(err)
	}

	db, err := backends.NewDatabase()

	if err != nil {
		log.Log.Fatal(err)
	}

	defer db.Close()

	if err = setupAuth(db); err != nil {
		log.Log.Fatal(err)
	}

	mux := mux.NewRouter()

	mux.PathPrefix("/auth").Handler(authManager.NewRouter())
	mux.Handle("/", authProtect(loggedIn, false))

	log.Log.Print("Starting HTTP server")

	log.Log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Cfg.Server.Hostname, config.Cfg.Server.Port), mux))
}

func loggedIn(w http.ResponseWriter, r *http.Request, user *datamodel.User) {
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(user.Email))
}

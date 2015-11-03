package main

import (
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"

	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/register"

	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
)

var (
	authManager = authboss.New()
	conf        = &Config{}
	log         = logrus.New()
)

func setupAuth(db *sql.DB) error {

	if err := InitializeCookieStore(conf); err != nil {
		return err
	}

	if err := InitializeSessionStore(conf); err != nil {
		return err
	}

	authManager.Storer = NewUserRepo(db)
	authManager.SessionStoreMaker = NewSessionStore
	authManager.CookieStoreMaker = NewCookieStore
	authManager.MountPath = "/auth"
	authManager.RootURL = "http://localhost:8080"
	authManager.LogWriter = logrus.StandardLogger().Out

	authManager.XSRFName = "csrf_token"
	authManager.XSRFMaker = func(_ http.ResponseWriter, req *http.Request) string {
		return nosurf.Token(req)
	}

	return authManager.Init()
}

func main() {
	log.Level = logrus.DebugLevel
	log.Print("Starting Reclus Issue Tracker...")

	if err := loadConfig(conf); err != nil {
		log.Fatal(err)
	}

	db, err := setupDatabase(conf)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = setupAuth(db); err != nil {
		log.Fatal(err)
	}

	mux := mux.NewRouter()

	mux.PathPrefix("/auth").Handler(authManager.NewRouter())

	http.ListenAndServe(":9090", mux)
}

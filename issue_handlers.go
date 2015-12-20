package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rdnelson/reclus/datamodel"
	"github.com/rdnelson/reclus/log"
)

func Issue(w http.ResponseWriter, r *http.Request, user *datamodel.User) {

	vars := mux.Vars(r)

	if err := json.NewEncoder(w).Encode(vars["id"]); err != nil {
		log.Log.Error(err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

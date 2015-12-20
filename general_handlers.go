package main

import (
	"encoding/json"
	"net/http"

	"github.com/rdnelson/reclus/datamodel"
	"github.com/rdnelson/reclus/log"
)

func Index(w http.ResponseWriter, r *http.Request, user *datamodel.User) {
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Log.Error(err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

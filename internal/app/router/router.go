package router

import (
	"github.com/dobuzora/bathrev/internal/app/models"
	"github.com/gorilla/mux"
)

func RegisterHandlers(vInfo models.VersionInfo) {
	r := mux.NewRouter()
	r.HandleFunc("/version").Methods("GET")

}

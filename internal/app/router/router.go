package router

import (
	"fmt"
	"github.com/dobuzora/bathrev/internal/app/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func RegisterHandlers(vInfo *models.VersionInfo) {
	r := mux.NewRouter()

	r.Methods("GET").Path("/version").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, vInfo)
		})

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

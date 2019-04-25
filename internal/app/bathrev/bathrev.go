package bathrev

import (
	"fmt"
	"github.com/dobuzora/bathrev/internal/app/database"
	"github.com/dobuzora/bathrev/internal/app/models"
	"github.com/dobuzora/bathrev/internal/app/router"
	"log"
	"net/http"
	"os"
)

var (
	Version   = "unkonwn"
	Commit    = "unkonwn"
	BuildData = "unkonwn"
)

func New() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	vInfo := &models.VersionInfo{Version: Version, Commit: Commit, BuildData: BuildData}
	router.RegisterHandlers(vInfo)
	db, err := database.New("sqlite3", "data", "admin", "admin", 10, true)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

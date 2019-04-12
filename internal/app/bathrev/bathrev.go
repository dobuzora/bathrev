package bathrev

import (
	"github.com/dobuzora/bathrev/internal/app/models"
	"github.com/dobuzora/bathrev/internal/app/router"
)

var (
	Version   = "unkonwn"
	Commit    = "unkonwn"
	BuildData = "unkonwn"
)

func New() {
	vInfo := &models.VersionInfo{Version: Version, Commit: Commit, BuildData: BuildData}
	router.RegisterHandlers(vInfo)

}

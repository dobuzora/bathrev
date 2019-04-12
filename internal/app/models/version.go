package models

type VersionInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildData string `json:"buildData"`
}

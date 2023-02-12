package version

import (
	"log"
	"runtime"
)

var (
	BuildDate string
	Version   string
	GitCommit string
	GitBranch string
	GoVersion = runtime.Version()
)

func init() {
	log.Println(
		"buildDate", BuildDate,
		"version", Version,
		"arch", runtime.GOARCH,
		"os", runtime.GOOS,
		"gitCommit", GitCommit,
		"gitBranch", GitBranch,
		"goVersion", GoVersion)
}

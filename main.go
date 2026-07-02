package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/ctrlpad/daemon/cmd"
	"github.com/ctrlpad/daemon/internal/info"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	log.SetDefault(log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		ReportCaller:    true,
	}))
	info.SetVersionInfo(version, commit, date)

	os.Exit(cmd.Run())
}

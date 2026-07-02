package info

import "github.com/charmbracelet/log"

func SetVersionInfo(version, commit, date string) {
	if commit != "none" {
		log.Infof("version: %s  commit: %s  built:  %s", version, commit, date)
	} else {
		log.Infof("Your running this application in development mode.")
	}
}

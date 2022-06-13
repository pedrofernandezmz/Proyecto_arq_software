package app

import (
	"os"
	log "github.com/sirupsen/logrus"
)

func init() {
	//INICIO DE LA APP
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("Starting system...")
}

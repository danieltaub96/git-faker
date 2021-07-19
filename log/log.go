package log

import log "github.com/sirupsen/logrus"
import . "github.com/danieltaub96/git-faker/object"

func InitLogger() {
	Formatter := new(log.TextFormatter)
	//Formatter.TimestampFormat = "28/02/2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

func CheckVerboseMode(u UserInput) {
	if u.Verbose {
		log.Info("Setting logger level to debug level")
		log.SetLevel(log.DebugLevel)
	}
}

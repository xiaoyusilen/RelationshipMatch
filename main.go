package main

import (
	"os"
	"runtime/debug"

	"RelationshipMatch/config"
	"RelationshipMatch/route"

	log "github.com/Sirupsen/logrus"
)

func initLogger(cfg *config.Config) {

	switch {
	case cfg.LogConfig.Formatter == config.LogTextFormatter:
		log.SetFormatter(&log.TextFormatter{})
	case cfg.LogConfig.Formatter == config.LogJSONFormatter:
		log.SetFormatter(&log.JSONFormatter{})
	}

	switch {
	case cfg.LogConfig.Output == config.LogConsoleOutput:
		log.SetOutput(os.Stdout)
	case cfg.LogConfig.Output == config.LogFileOutput:
		f, err := os.OpenFile(cfg.LogConfig.FilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Panicf("error opening log file: %v", err)
		}

		log.SetOutput(f)
	}

	switch {
	case cfg.LogConfig.Level == config.LogPanicLevel:
		log.SetLevel(log.PanicLevel)
	case cfg.LogConfig.Level == config.LogFatalLevel:
		log.SetLevel(log.FatalLevel)
	case cfg.LogConfig.Level == config.LogErrorLevel:
		log.SetLevel(log.ErrorLevel)
	case cfg.LogConfig.Level == config.LogWarnLevel:
		log.SetLevel(log.WarnLevel)
	case cfg.LogConfig.Level == config.LogInfoLevel:
		log.SetLevel(log.InfoLevel)
	case cfg.LogConfig.Level == config.LogDebugLevel:
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	// parse config
	cfg := config.ParseFromFlags()

	// init log
	initLogger(cfg)

	// init Http service
	r := route.HandleRest(cfg)

	r.Router.Run(":9000")

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}

		// close PG DB
		r.PG.Close()
	}()
}

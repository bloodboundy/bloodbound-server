package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func setupLogger() {
	var level log.Level
	err := level.UnmarshalText([]byte(*LEVEL))
	if err != nil {
		log.Panicf("unaccptable level: %v, accepts: trace debug info warning error fatal panic", *LEVEL)
	}
	log.SetLevel(level)

	out := os.Stderr
	if *OUTPUT != "" {
		file, err := os.OpenFile(*OUTPUT, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Panicf("failed to open log output file: %v", err)
		}
		out = file
	}
	log.SetOutput(out)
}

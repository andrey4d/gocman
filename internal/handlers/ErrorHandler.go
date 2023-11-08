package handlers

import (
	log "github.com/sirupsen/logrus"
)

func ErrorHandler(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Fatalf("ERROR: %v", err)
	}

}

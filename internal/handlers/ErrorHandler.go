package handlers

import (
	log "github.com/sirupsen/logrus"
)

func ErrorHandlerPanicWithMessage(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Fatalf("ERROR: %v", err)
	}

}
func ErrorHandlerReturn(err error, msg string) error {
	if err != nil {
		log.Println(msg)
		return err
	}
	return nil
}

func ErrorHandlerLog(msg string) {
	log.Error(msg)
}

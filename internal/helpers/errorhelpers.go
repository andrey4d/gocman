/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package helpers

import (
	log "github.com/sirupsen/logrus"
)

func CheckError(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Fatalf("ERROR: %v", err)
	}

}
func ErrorHelperReturn(err error, msg string) error {
	if err != nil {
		log.Println(msg)
		return err
	}
	return nil
}

func ErrorHelperLog(msg string) {
	log.Error(msg)
}

func FatalHelperLog(msg string) {
	log.Fatal(msg)
}

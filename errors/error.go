package errors

import "log"

func LogAnyErrorAndExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

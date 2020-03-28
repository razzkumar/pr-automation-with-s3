package logger

import (
	"fmt"
	"log"
)


func FailOnNoFlag(msg string) {
	panic(fmt.Sprintf("%s", msg))
}

func Info(msg string) {
	log.Panic(msg)
}

func FailOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

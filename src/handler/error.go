package handler

import "merger/src/logger"

func FailOnError(err error, message string) {
	if err != nil {
		logger.Error.Println(err.Error())
		logger.Info.Println(message)
		panic(err)
	}
}
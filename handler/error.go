package handler

import "merger/logger"

func FailOnError(err error, message string) {
	if err != nil {
		logger.Error.Println(err.Error())
		logger.Info.Println(message)
		panic(err)
	}
}
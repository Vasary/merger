package main

import (
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"merger/src/logger"
	"merger/src/processor/prepare"
	"merger/src/processor/rematcher"
	"merger/src/processor/score"
	"merger/src/service/database"
	"merger/src/service/sort"
	"merger/src/service/user/provider"
	"os"
	"os/signal"
	"syscall"
)

var env = []string{
	"LOG_PATH",
	"DATABASE_HOST",
	"DATABASE_USER",
	"DATABASE_PASSWORD",
	"DATABASE_NAME",
}

func init() {
	for _, v := range env {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("Environment variable %v has unexpacted value %v", v, os.Getenv(v)))
		}
	}
}

func main() {
	logger.Info("Starting")

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		users := provider.GetUsers(database.GetConnection())
		logger.Debug(fmt.Sprintf("Found %d records for add up", len(users)))

		if len(users) > 0 {
			bar := pb.StartNew(len(users))

			for _, row := range users {
				rematcher.Rematch(sort.Sort(score.Score(prepare.Prepare(row, database.GetConnection()))), database.GetConnection())

				bar.Increment()
			}

			bar.Finish()
		}

		defer database.GetConnection().Close()

		logger.Info("Completed")

		os.Exit(0)
	}()

	<-stop

	logger.Info("Shutdown")
}

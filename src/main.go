package main

import (
	"merger/src/logger"
	"merger/src/service/database"
	"merger/src/service/user/provider"
	"merger/src/processor/prepare"
	"merger/src/processor/score"
	"merger/src/processor/rematcher"
	"merger/src/service/sort"
	"fmt"
)

func main() {
	logger.Info("Starting")

	users := provider.GetUsers(database.GetConnection())
	logger.Info(fmt.Sprintf("Found %d records for add up", len(users)))

	for _, row := range users {
		rematcher.Rematch(sort.Sort(score.Score(prepare.Prepare(row, database.GetConnection()))))
	}

	defer database.GetConnection().Close()
}

package main

import (
	"fmt"
	"merger/src/logger"
	"merger/src/processor/prepare"
	"merger/src/processor/rematcher"
	"merger/src/processor/score"
	"merger/src/service/database"
	"merger/src/service/sort"
	"merger/src/service/user/provider"
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

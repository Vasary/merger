package main

import (
	"usersRematcher/logger"
	"usersRematcher/service/database"
	"usersRematcher/service/user/provider"
	"usersRematcher/processor/prepare"
	"usersRematcher/processor/score"
	"usersRematcher/processor/rematcher"
	"fmt"
)

func main() {
	logger.Info.Println("Starting")

	users := provider.GetUsers()
	logger.Info.Println(fmt.Sprintf("Found %d records for addup", len(users)))

	for _, row := range users {
		rematcher.Rematch(score.Score(prepare.Prepare(row)))
	}

	defer database.GetConnection().Close()
}

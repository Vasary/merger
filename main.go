package main

import (
	"merger/logger"
	"merger/service/database"
	"merger/service/user/provider"
	"merger/processor/prepare"
	"merger/processor/score"
	"merger/processor/rematcher"
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

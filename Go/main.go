package main

import (
	"fmt"
	"myapp/controller"
	"myapp/database"
)

func main() {

	db, connError := database.DatabaseConn()

	if connError != nil {
		fmt.Println("Issue finding database. Closing...")
		return
	}

	defer db.Close()

	controller.Run(db)

}

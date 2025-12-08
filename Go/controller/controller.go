package controller

import (
	"database/sql"
	"fmt"
	"myapp/ui"

	_ "github.com/go-sql-driver/mysql"
)

func Run(db *sql.DB) {

	for {

		menuOption := ui.GetMenuOption()

		switch menuOption {
		case 1:
			listRecords(db)
		case 2:
			createRecords(db)
		case 3:
			viewRecords(db)
		case 4:
			removeRecords(db)
		default:
			return
		}

	}

}

func listRecords(db *sql.DB) {
	fmt.Println("List")
}

func createRecords(db *sql.DB) {
	fmt.Println("Create")
}

func viewRecords(db *sql.DB) {
	fmt.Println("View")
}

func removeRecords(db *sql.DB) {
	fmt.Println("Remove")
}

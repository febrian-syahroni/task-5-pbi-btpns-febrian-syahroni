package main

import (
	"task-5-pbi-btpns-febrian-syahroni/database"
	"task-5-pbi-btpns-febrian-syahroni/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}

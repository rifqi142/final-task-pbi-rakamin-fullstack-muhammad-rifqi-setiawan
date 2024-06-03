package main

import (
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/database"
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}
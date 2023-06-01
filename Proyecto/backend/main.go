package main

import (
	"Proyecto/app"
	"Proyecto/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}

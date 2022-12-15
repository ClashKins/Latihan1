package main

import (
	"LATIHAN1/database"
	"LATIHAN1/router"
)

func main() {
	database.StarDB()
	r := router.StartApp()
	r.Run(":8080")
}
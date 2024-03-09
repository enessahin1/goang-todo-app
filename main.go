package main

import (
	"todo_app/src/config"
	"todo_app/src/routes"
)

func main() {
	defer config.DisconnectDB()

	//run all routes
	routes.Routes()
}

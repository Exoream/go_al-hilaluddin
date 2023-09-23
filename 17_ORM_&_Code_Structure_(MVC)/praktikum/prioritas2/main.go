package main

import (
	"db_API/prioritas2/config"
	"db_API/prioritas2/routes"
)
func main() {
	config.Init()
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

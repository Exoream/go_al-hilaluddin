package main

import (
	"db_API/config"
	"db_API/middlewares"
	"db_API/routes"
)
func main() {
	config.Init()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

package main

import (
	"db_API/prioritas1/nomor2/config"
	"db_API/prioritas1/nomor2/route"
)
func main() {
	config.Init()
	e := route.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

package main

import (
	"clean/api/config"
	"clean/api/router"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var cfg = config.InitConfig()

	dbMysql := config.InitMysqlConn(cfg)
	router.NewRoute(e, dbMysql)

	config.Migrate(dbMysql)
	
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}

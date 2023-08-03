package main

import (
	"saichudin/kaspin/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
    router.Route(e)
    e.Logger.Fatal(e.Start(":1323"))
}
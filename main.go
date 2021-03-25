package main

import (
	"hapoon/go-api-template/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.Route(e)

	e.Logger.Fatal(e.Start(":1323"))
}

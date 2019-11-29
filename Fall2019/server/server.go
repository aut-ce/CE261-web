package main

import (
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// try to allow cors from anywhere
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The hallows.ir server is up")
	})

	e.GET("/chart", func(c echo.Context) error {
		return c.File("./data/chart.json")
	})
	e.GET("/occasion", func(c echo.Context) error {
		return c.File("./data/occasion.json")
	})
	e.GET("/house/:id", func(c echo.Context) error {
		return c.File("./data/house.json")
	})
	e.GET("/mags", func(c echo.Context) error {
		return c.File("./data/mags.json")
	})

	e.Logger.Fatal(e.Start(":80"))
}

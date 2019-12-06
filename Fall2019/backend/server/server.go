package main

import (
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"context"
)

func main() {
	e := echo.New()

	mongoClient := GetClient()
	err := mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

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
		house := GetHouse(mongoClient, c.Param("id"))
		if house == nil {
			return c.JSON(http.StatusNotFound, "")
		}
		return c.JSON(http.StatusNotFound, house)
	})

	e.GET("/mags", func(c echo.Context) error {
		return c.File("./data/mags.json")
	})

	e.Logger.Fatal(e.Start(":80"))
}

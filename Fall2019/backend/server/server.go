package main

import (
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"context"
)

const (
	// URL
	BaseUrl = "http://hallows.ir"
	//BaseUrl = "http://localhost"
	// Magazine
	MagazineTitle = "کیلید مگ را بخوانید"

	// Occasion
	OccasionTitle      = "اکازیون های امروز کیلید"
	LimitOccasion      = 4
	EndOccasion        = 8
	OccasionActionText = "اکازیون های بیشتر در کیلید"

	//Chart
	SelectedColor = "#ffdf00"
	OtherColor    = "#ff5f00"
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
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type,Cache-Control,Pragma")
			c.Response().Header().Set("Access-Control-Max-Age", "2592000")
			return nil
		}

	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The hallows.ir server is up")
	})

	e.GET("/chart", func(c echo.Context) error {
		chart := GetAllChartHouse(mongoClient)
		chart.SelectedColor = SelectedColor
		chart.OtherColor = OtherColor

		return c.JSON(http.StatusOK, chart)
	})

	e.GET("/occasion", func(c echo.Context) error {
		occasion := GetOccasion(mongoClient, LimitOccasion, 0)
		occasion.Section = OccasionTitle
		occasion.Action = (*struct {
			Text string `json:"text"`
			URL  string `json:"url"`
		})(&struct {
			Text string
			URL  string
		}{
			Text: OccasionActionText,
			URL:  BaseUrl + "/occasion/" + strconv.Itoa(LimitOccasion),
		})
		return c.JSON(http.StatusOK, occasion)
	})

	e.GET("/occasion/:skipped", func(c echo.Context) error {
		skipped, err := strconv.Atoi(c.Param("skipped"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "")
		}

		occasion := GetOccasion(mongoClient, LimitOccasion, skipped)
		occasion.Section = OccasionTitle

		if skipped+LimitOccasion < EndOccasion {
			occasion.Action = (*struct {
				Text string `json:"text"`
				URL  string `json:"url"`
			})(&struct {
				Text string
				URL  string
			}{
				Text: OccasionActionText,
				URL:  BaseUrl + "/occasion/" + strconv.Itoa(LimitOccasion+skipped),
			})
		}
		return c.JSON(http.StatusOK, occasion)
	})

	e.GET("/house/:id", func(c echo.Context) error {
		house := GetHouse(mongoClient, c.Param("id"))
		if house == nil {
			return c.JSON(http.StatusNotFound, "")
		}
		return c.JSON(http.StatusOK, house)
	})

	e.GET("/mags", func(c echo.Context) error {
		mags := GetAllMagazine(mongoClient)
		mags.Section = MagazineTitle
		return c.JSON(http.StatusOK, mags)
	})

	e.Logger.Fatal(e.Start(":80"))
}

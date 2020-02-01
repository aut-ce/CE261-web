package main

import (
	"github.com/aut-ce/Web101/mongo"
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
	MagazineTitle = "مطالب مجیله مگ"

	// Occasion
	OccasionTitle      = "خونه های خالی امروز کیلید"
	LimitOccasion      = 4
	EndOccasion        = 8
	OccasionActionText = "خونه‌های بیشتر"

	//Chart
	SelectedColor = "#00dfff"
	OtherColor    = "#ff005f"
)

func main() {
	e := echo.New()

	mongoClient := mongo.GetClient()
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
		chart := mongo.GetAllChartHouse(mongoClient)
		chart.SelectedColor = SelectedColor
		chart.OtherColor = OtherColor

		return c.JSON(http.StatusOK, chart)
	})

	e.GET("/occasion", func(c echo.Context) error {
		occasion := mongo.GetOccasion(mongoClient, LimitOccasion, 0)
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

		occasion := mongo.GetOccasion(mongoClient, LimitOccasion, skipped)
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
		house := mongo.GetHouse(mongoClient, c.Param("id"))
		if house == nil {
			return c.JSON(http.StatusNotFound, "")
		}
		return c.JSON(http.StatusOK, house)
	})

	e.GET("/mags", func(c echo.Context) error {
		mags := mongo.GetAllMagazine(mongoClient)
		mags.Section = MagazineTitle
		return c.JSON(http.StatusOK, mags)
	})

	e.Logger.Fatal(e.Start(":80"))
}

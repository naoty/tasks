package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/naoty/tasks/backend/model"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/statuses", getStatuses)
	e.Logger.Fatal(e.Start(":1323"))
}

func getStatuses(c echo.Context) error {
	statuses := []*model.Status{
		&model.Status{ID: "1", Name: "TODO"},
		&model.Status{ID: "2", Name: "DOING"},
		&model.Status{ID: "3", Name: "DONE"},
	}
	return c.JSON(http.StatusOK, statuses)
}

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
	e.Use(middleware.CORS())
	e.GET("/statuses", getStatuses)
	e.Logger.Fatal(e.Start(":1323"))
}

func getStatuses(c echo.Context) error {
	statuses := []*model.Status{
		&model.Status{ID: "1", Name: "TODO", TaskIDs: []string{}},
		&model.Status{ID: "2", Name: "DOING", TaskIDs: []string{}},
		&model.Status{ID: "3", Name: "DONE", TaskIDs: []string{}},
	}
	return c.JSON(http.StatusOK, statuses)
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/naoty/tasks/backend/model"
)

var dsn = fmt.Sprintf(
	"%s:%s@tcp(%s:3306)/%s",
	os.Getenv("DATABASE_USER"),
	os.Getenv("DATABASE_PASSWORD"),
	os.Getenv("DATABASE_HOST"),
	os.Getenv("DATABASE_NAME"),
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.GET("/statuses", getStatuses)
	e.Logger.Fatal(e.Start(":1323"))
}

func getStatuses(c echo.Context) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to connect db")
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error: %v", err))
	}

	statuses := []*model.Status{
		&model.Status{ID: "1", Name: "TODO", TaskIDs: []string{}},
		&model.Status{ID: "2", Name: "DOING", TaskIDs: []string{}},
		&model.Status{ID: "3", Name: "DONE", TaskIDs: []string{}},
	}
	return c.JSON(http.StatusOK, statuses)
}

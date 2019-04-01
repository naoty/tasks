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
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.GET("/statuses", getStatuses)
	e.Logger.Fatal(e.Start(":1323"))
}

func getStatuses(c echo.Context) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM statuses")
	if err != nil {
		return err
	}
	defer rows.Close()

	statuses := []model.Status{}
	for rows.Next() {
		status := model.Status{TaskIDs: []string{}}
		err := rows.Scan(&status.StatusID, &status.Name)
		if err != nil {
			return err
		}
		statuses = append(statuses, status)
	}

	return c.JSON(http.StatusOK, statuses)
}

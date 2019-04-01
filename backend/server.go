package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/naoty/tasks/backend/model"
)

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()
	e.Debug = config.Debug
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(CustomContextMiddleware)
	e.Use(ConfigMiddleware(config))
	e.GET("/statuses", getStatuses)
	e.Logger.Fatal(e.Start(":1323"))
}

func getStatuses(c echo.Context) error {
	cc := c.(*CustomContext)
	db, err := sql.Open("mysql", cc.GetDSN())
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

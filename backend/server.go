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
	e.Use(databaseMiddleware)

	e.GET("/statuses", getStatuses)

	e.Logger.Fatal(e.Start(":1323"))
}

func databaseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)
		db, err := sql.Open("mysql", cc.GetDSN())
		if err != nil {
			return err
		}
		defer db.Close()

		cc.DB = db
		return next(cc)
	}
}

func getStatuses(c echo.Context) error {
	cc := c.(*CustomContext)
	rows, err := cc.Query("SELECT * FROM statuses LEFT OUTER JOIN tasks USING (status_id)")
	if err != nil {
		return err
	}
	defer rows.Close()

	type result struct {
		statusID string
		name     string
		taskID   sql.NullString
		title    sql.NullString
	}

	statusMap := map[string]model.Status{}
	for rows.Next() {
		var result result
		err := rows.Scan(&result.statusID, &result.name, &result.taskID, &result.title)
		if err != nil {
			return err
		}

		status, ok := statusMap[result.statusID]
		if !ok {
			status = model.Status{StatusID: result.statusID, Name: result.name, Tasks: []model.Task{}}
		}

		if result.taskID.Valid {
			task := model.Task{TaskID: result.taskID.String, Title: result.title.String, StatusID: result.statusID}
			status.Tasks = append(status.Tasks, task)
		}

		statusMap[result.statusID] = status
	}

	statuses := []model.Status{}
	for _, status := range statusMap {
		statuses = append(statuses, status)
	}

	return c.JSON(http.StatusOK, statuses)
}

package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/naoty/tasks/backend/config"
	"github.com/naoty/tasks/backend/handler"
)

func main() {
	var env config.Env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()
	e.Debug = env.Debug

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(handler.CustomContextMiddleware)
	e.Use(handler.EnvMiddleware(env))
	e.Use(handler.DatabaseMiddleware)

	e.GET("/statuses", handler.GetStatuses)
	e.POST("/tasks", handler.PostTasks)

	e.Logger.Fatal(e.Start(":1323"))
}

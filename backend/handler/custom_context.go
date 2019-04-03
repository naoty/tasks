package handler

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/naoty/tasks/backend/config"
)

// CustomContext is the extension of echo.Context.
type CustomContext struct {
	echo.Context
	config.Env
	*sql.DB
}

// CustomContextMiddleware is a middleware to extend echo.Context.
func CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{Context: c}
		return next(cc)
	}
}

// EnvMiddleware is a middleware to inject Config to CustomContext.
func EnvMiddleware(env config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*CustomContext)
			cc.Env = env
			return next(cc)
		}
	}
}

// DatabaseMiddleware is a middleware to manage connections to database.
func DatabaseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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

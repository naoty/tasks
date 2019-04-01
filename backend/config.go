package main

import (
	"fmt"

	"github.com/labstack/echo"
)

// Config represents configurations of this application.
type Config struct {
	DatabaseHost     string `required:"true" split_words:"true"`
	DatabaseName     string `required:"true" split_words:"true"`
	DatabasePassword string `required:"true" split_words:"true"`
	DatabaseUser     string `required:"true" split_words:"true"`
	Debug            bool   `default:"false"`
}

// GetDSN returns the Data Source Name of database.
func (config Config) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabaseName,
	)
}

// ConfigMiddleware is a middleware to inject Config to CustomContext.
func ConfigMiddleware(config Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*CustomContext)
			cc.Config = config
			return next(cc)
		}
	}
}

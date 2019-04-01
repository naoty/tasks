package main

import "github.com/labstack/echo"

// CustomContext is the extension of echo.Context.
type CustomContext struct {
	echo.Context
	Config
}

// CustomContextMiddleware is a middleware to extend echo.Context.
func CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{Context: c}
		return next(cc)
	}
}

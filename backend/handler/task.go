package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// PostTasks is a handler for `POST /tasks`.
func PostTasks(c echo.Context) error {
	type BodyTask struct {
		Title string `json:"title"`
	}

	type Body struct {
		Task BodyTask `json:"task"`
	}

	body := new(Body)
	if err := c.Bind(body); err != nil {
		return err
	}

	cc := c.(*CustomContext)
	rows, err := cc.Query("SELECT status_id FROM statuses ORDER BY position ASC LIMIT 1")
	if err != nil {
		return err
	}

	rows.Next()
	var initialStatusID int
	err = rows.Scan(&initialStatusID)
	if err != nil {
		return err
	}

	_, err = cc.Exec("INSERT INTO tasks (title, status_id) VALUES (?, ?)", body.Task.Title, initialStatusID)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

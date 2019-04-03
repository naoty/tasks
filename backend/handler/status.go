package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoty/tasks/backend/model"
)

// GetStatuses is a handler for `GET /statuses`.
func GetStatuses(c echo.Context) error {
	cc := c.(*CustomContext)
	rows, err := cc.Query("SELECT * FROM statuses LEFT OUTER JOIN tasks USING (status_id) ORDER BY position ASC")
	if err != nil {
		return err
	}
	defer rows.Close()

	type result struct {
		statusID string
		name     string
		position int
		taskID   sql.NullString
		title    sql.NullString
	}

	statusMap := map[string]model.Status{}
	for rows.Next() {
		var result result
		err := rows.Scan(
			&result.statusID,
			&result.name,
			&result.position,
			&result.taskID,
			&result.title,
		)
		if err != nil {
			return err
		}

		status, ok := statusMap[result.statusID]
		if !ok {
			status = model.Status{
				StatusID: result.statusID,
				Name:     result.name,
				Position: result.position,
				Tasks:    []model.Task{},
			}
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

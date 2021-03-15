package manager

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func RegisterHandler(e *echo.Echo, service Service) {
	e.POST("/api/task", newTask(service))
	e.PUT("/api/task/:taskId", taskDone(service))
	e.DELETE("/api/task/:taskId", deleteTask(service))
	e.GET("/api/task/:taskId", findById(service))
	e.GET("/api/task", findAll(service))
}

func newTask(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := c.FormValue("name")
		err := service.newTask(task)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, "created")
	}
}

func taskDone(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId := c.Param("taskId")
		n, err := service.taskDone(taskId)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		if n == 0 {
			return c.JSON(http.StatusNotFound, "item not found")
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func deleteTask(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId := c.Param("taskId")
		n, err := service.delete(taskId)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		if n == 0 {
			return c.JSON(http.StatusNotFound, "item not found")
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func findById(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId := c.Param("taskId")
		task, err := service.findById(taskId)

		switch err {
		case sql.ErrNoRows:
			return c.JSON(http.StatusNotFound, "no result")
		case nil:
			return c.JSON(http.StatusOK, task)
		default:
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
}

func findAll(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		todoList, err := service.findAll()
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, todoList)
	}
}

package handlers

import (
	"github.com/labstack/echo/v4"
	"go.mod/internal/taskService"
	"net/http"
)

type TaskHandler struct {
	service *taskService.Service
}

func NewTaskHandler(service *taskService.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTask(c echo.Context) error {
	var task taskService.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	if err := h.service.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create task"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) PatchTask(c echo.Context) error {
	id := c.Param("id")
	var updateTask taskService.Task

	if err := c.Bind(&updateTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	task, err := h.service.UpdateTask(id, &updateTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

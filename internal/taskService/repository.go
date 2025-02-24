package taskService

import "go.mod/internal/database"

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := database.DB.Find(&tasks)
	return tasks, result.Error
}

func CreateTask(task *Task) error {
	result := database.DB.Create(task)
	return result.Error
}

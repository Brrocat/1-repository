package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Message string `json:"message"`
	IsDone  bool   `json:"isDone"`
}

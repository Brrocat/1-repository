package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name    string `json:"name"`
	Message string `json:"message"`
	IsDone  bool   `json:"is_done"`
	UserID  uint   `json:"user_id"`
}

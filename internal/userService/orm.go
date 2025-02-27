package userService

import (
	"go.mod/internal/taskService"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks" gorm:"foreignkey:UserID"`
}

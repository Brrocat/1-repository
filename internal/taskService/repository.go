package taskService

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

func (r *Repository) CreateTask(task *Task) error {
	result := r.db.Create(task)

	return result.Error
}

func (r *Repository) UpdateTask(id string, task *Task) (*Task, error) {
	var existingTask Task
	result := r.db.First(&existingTask, id)
	if result.Error != nil {
		return nil, result.Error
	}

	existingTask.Name = task.Name
	existingTask.Message = task.Message
	existingTask.IsDone = task.IsDone

	result = r.db.Save(&existingTask)
	return &existingTask, result.Error
}

func (r *Repository) DeleteTask(id string) error {
	result := r.db.Delete(&Task{}, "id = ?", id)
	return result.Error
}

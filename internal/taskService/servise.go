package taskService

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *Service) CreateTask(task *Task) error {
	return s.repo.CreateTask(task)
}

func (s *Service) UpdateTask(id string, task *Task) (*Task, error) {
	return s.repo.UpdateTask(id, task)
}

func (s *Service) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

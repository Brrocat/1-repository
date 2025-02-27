package userService

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUser()
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) UpdateUser(id string, user *User) (*User, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *Service) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

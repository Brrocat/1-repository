package userService

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (h *Repository) GetAllUser() ([]User, error) {
	var users []User
	result := h.db.Find(&users)
	return users, result.Error
}

func (h *Repository) CreateUser(user *User) error {
	result := h.db.Create(user)
	return result.Error
}

func (h *Repository) UpdateUser(id string, user *User) (*User, error) {
	var existingUser User
	result := h.db.First(&existingUser, id)
	if result.Error != nil {
		return nil, result.Error
	}

	existingUser.Email = user.Email
	existingUser.Password = user.Password

	result = h.db.Save(user)
	return &existingUser, result.Error
}

func (h *Repository) DeleteUser(id string) error {
	result := h.db.Delete(&User{}, "id = ?", id)
	return result.Error
}

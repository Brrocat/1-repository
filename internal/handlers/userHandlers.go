package handlers

import (
	"github.com/labstack/echo/v4"
	"go.mod/internal/userService"
	"net/http"
)

type UserHandler struct {
	service *userService.Service
}

func NewUserHandler(service *userService.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) PostUser(c echo.Context) error {
	var user userService.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	if err := h.service.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) PatchUserId(c echo.Context) error {
	id := c.Param("id")
	var updateUser userService.User

	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	user, err := h.service.UpdateUser(id, &updateUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	return c.NoContent(http.StatusNoContent)
}

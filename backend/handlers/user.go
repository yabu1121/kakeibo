package handlers

import (
	"kakeibo-backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandlers struct {
	DB *gorm.DB
}

// CREATE
func (h *UserHandlers) CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := h.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, user)
}

// GET
func (h *UserHandlers) GetUsers(c echo.Context) error {
	user := []models.User{}
	if err := h.DB.Find(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, user)
}

// GET by id
func (h *UserHandlers) GetUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	user := models.User{}
	if err := h.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, user)
}

// GET by name
func (h *UserHandlers) SearchUser(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}
	users := []models.User{}
	if err := h.DB.Where("name LIKE ?", "%"+name+"%").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, users)
}

// GET with pagination
func (h *UserHandlers) GetUserWithPagination(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "page is required"})
	}
	if limit == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "limit is required"})
	}
	offset := (page - 1) * limit
	users := []models.User{}
	if err := h.DB.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, users)
}

// UPDATE(PUT)
func (h *UserHandlers) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	type UpdateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
	}

	if err := h.DB.Model(&models.User{}).Where("id=?", id).Updates(req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, req)
}

// DELETE
func (h *UserHandlers) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	if err := h.DB.Where("id=?", id).Delete(&models.User{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, nil)
}

package handlers

import (
	"kakeibo-backend/models"
	"net/http"

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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}
	return c.JSON(http.StatusOK, user)
}


// GET
func (h *UserHandlers) GetUsers(c echo.Context) error {
	user := []models.User{}
	if err := h.DB.Find(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
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

	if err := h.DB.Model(&models.User{}).Where("id=?",id).Updates(req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"internal server error"})
	}
	return c.JSON(http.StatusOK, req)
}



// DELETE
func (h *UserHandlers) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error":"id is required"})
	}
	
	if err := h.DB.Where("id=?",id).Delete(&models.User{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"internal server error"})
	}
	return c.JSON(http.StatusOK, nil)
}
package handlers

import (
	"kakeibo-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryHandlers struct {
	DB *gorm.DB
}

func (h *CategoryHandlers) CreateCategory(c echo.Context) error {
	category := models.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, category)
}

func (h *CategoryHandlers) GetCategory(c echo.Context) error {
	categories := []models.Category{}
	if err := h.DB.Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandlers) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	type UpdateCategoryRequest struct {
		Name string `json:"name" gorm:"not null;unique"`
	}
	var req UpdateCategoryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.DB.Model(&models.Category{}).Where("id = ?", id).Updates(&req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *CategoryHandlers) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	if err := h.DB.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.NoContent(http.StatusNoContent)
}

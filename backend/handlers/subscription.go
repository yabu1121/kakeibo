package handlers

import (
	"kakeibo-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SubscriptionHandlers struct {
	DB *gorm.DB
}

func (h *SubscriptionHandlers) CreateSubscription(c echo.Context) error {
	subscription := models.Subscription{}
	if err := c.Bind(&subscription); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.DB.Create(&subscription).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}
	return c.JSON(http.StatusOK, subscription)
}
func (h *SubscriptionHandlers) GetSubscription(c echo.Context) error {
	subscription := []models.Subscription{}
	if err := h.DB.Find(&subscription).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}
	return c.JSON(http.StatusOK, subscription)
}

func (h *SubscriptionHandlers) UpdateSubscription(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	type UpdateSubscriptionRequest struct {
		ProductName string `json:"product_name"`
		CategoryID  uint   `json:"category_id"`
		Amount      int    `json:"amount"`
		Frequency   string `json:"frequency"`
	}
	var req UpdateSubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.DB.Model(&models.Subscription{}).Where("id=?", id).Updates(req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, req)
}

func (h *SubscriptionHandlers) DeleteSubscription(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	if err := h.DB.Where("id=?", id).Delete(&models.Subscription{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.NoContent(http.StatusNoContent)
}

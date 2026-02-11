package handlers

import (
	"kakeibo-backend/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ExpenseHandlers struct {
	DB *gorm.DB
}

// CREATE
func (h *ExpenseHandlers) CreateExpense (c echo.Context) error {
	expense := models.Expense{}
	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.DB.Create(&expense).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, expense)
}



// READ
func (h *ExpenseHandlers) GetExpense (c echo.Context) error {
	expense := []models.Expense{}
	// DB.preload("Category")とすればIdからそのままgetできる。
	if err := h.DB.Find(&expense).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, expense)
}


// UPDATE
func (h *ExpenseHandlers) UpdateExpense (c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	// gorm部分は記述しなくていい
	type UpdateExpenseRequest struct {
		Amount     int       `json:"amount"`
		Date       time.Time `json:"date"`        
		Memo       string    `json:"memo"`                            
		CategoryID uint      `json:"category_id"` 
	}
	var req UpdateExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// Modelチェーン
	if err := h.DB.Where("id = ?", id).Model(&models.Expense{}).Updates(req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"internal server error"})
	}
	return c.JSON(http.StatusOK, req)
}



// DELETE
func (h *ExpenseHandlers) DeleteExpense (c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}
	// deleteの中身
	if err := h.DB.Where("id = ?", id).Delete(&models.Expense{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.NoContent(http.StatusNoContent)
}
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
func (h *ExpenseHandlers) CreateExpense(c echo.Context) error {
	type CreateExpenseRequest struct {
		UserID     uint   `json:"user_id"`
		Amount     int    `json:"amount"`
		Date       string `json:"date"` // 文字列で受け取る
		Memo       string `json:"memo"`
		CategoryID uint   `json:"category_id"`
	}

	req := CreateExpenseRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// 日付パース (YYYY-MM-DD)
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		// RFC3339も試す
		parsedDate, err = time.Parse(time.RFC3339, req.Date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid date format. use YYYY-MM-DD or RFC3339"})
		}
	}

	expense := models.Expense{
		UserID:     req.UserID,
		Amount:     req.Amount,
		Date:       parsedDate,
		Memo:       req.Memo,
		CategoryID: req.CategoryID,
	}

	if err := h.DB.Create(&expense).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, expense)
}

// READ
func (h *ExpenseHandlers) GetExpense(c echo.Context) error {
	expense := []models.Expense{}
	// DB.preload("Category")とすればIdからそのままgetできる。
	if err := h.DB.Find(&expense).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, expense)
}

// UPDATE
func (h *ExpenseHandlers) UpdateExpense(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	type UpdateExpenseRequest struct {
		Amount     int    `json:"amount"`
		Date       string `json:"date"` // 文字列で受け取る
		Memo       string `json:"memo"`
		CategoryID uint   `json:"category_id"`
	}
	var req UpdateExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	updates := map[string]interface{}{}
	if req.Amount != 0 {
		updates["amount"] = req.Amount
	}
	if req.Date != "" {
		parsedDate, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			parsedDate, err = time.Parse(time.RFC3339, req.Date)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid date format"})
			}
		}
		updates["date"] = parsedDate
	}
	if req.Memo != "" {
		updates["memo"] = req.Memo
	}
	if req.CategoryID != 0 {
		updates["category_id"] = req.CategoryID
	}

	// Modelチェーン
	if err := h.DB.Model(&models.Expense{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, req)
}

// DELETE
func (h *ExpenseHandlers) DeleteExpense(c echo.Context) error {
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


// 日にちごとの支出を取得
func (h *ExpenseHandlers) GetExpenseByDate(c echo.Context) error {
    dateStr := c.Param("date") // 例: "2026-02-12"

    // 1. 文字列を time.Time 型に変換（バリデーションを兼ねる）
    layout := "2006-01-02"
    t, err := time.Parse(layout, dateStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid date format. use YYYY-MM-DD"})
    }

    // 2. その日の開始時刻と終了時刻を計算
    startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
    endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

    var expenses []models.Expense
    
    // 3. 範囲検索（BETWEEN）を使う
    // GORMではこのように書くと、時間のずれを気にせず「その日1日分」を確実に取れます
    if err := h.DB.Where("date BETWEEN ? AND ?", startOfDay, endOfDay).Find(&expenses).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
    }

    return c.JSON(http.StatusOK, expenses)
}
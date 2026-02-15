package handlers

import (
	"kakeibo-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}


// CREATE
func (h *UserHandler) CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}



// GET
func (h *UserHandler) GetAllUser(c echo.Context) error {
	user := []models.User{}
	if err := h.DB.Find(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}



// GET BY ID
func (h *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	if err := h.DB.First(&user, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}




// UPDATE
func (h *UserHandler) UpdateUser(c echo.Context) error {
    id := c.Param("id")
    var user models.User
    if err := h.DB.First(&user, "id = ?", id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, "User not found")
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if err := h.DB.Save(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, user)
}




// DELETE
func (h *UserHandler) DeleteUserById(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

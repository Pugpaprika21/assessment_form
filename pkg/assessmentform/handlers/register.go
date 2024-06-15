package handlers

import (
	"net/http"

	"github.com/Pugpaprika21/internal/dto"
	"github.com/Pugpaprika21/pkg/assessmentform/models"
	"github.com/Pugpaprika21/pkg/assessmentform/server"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type register struct {
	query *gorm.DB
}

func NewRegister(server *server.EchoServerEnvironment) *register {
	return &register{
		query: server.Connect.ORM,
	}
}

func (r *register) FormRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{
		"title": "ลงทะเบียน",
	})
}

func (r *register) RegisterSave(c echo.Context) error {
	if c.Request().Header.Get("X-Requested-With") != "XMLHttpRequest" {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    "ไม่มีข้อมูลที่ถูกส่งมา",
			StatusBool: false,
		})
	}

	var userEmailCount int64
	var userModel models.User
	r.query.Model(&userModel).Where("email = ?", c.FormValue("email")).Count(&userEmailCount)
	if userEmailCount != 0 {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.FailedMessageSaveTH + " มีอีเมลที่ถูกใช้งานเเล้ว",
			StatusBool: false,
		})
	}

	userModel.Fullname = c.FormValue("fullname")
	userModel.Email = c.FormValue("email")
	userModel.Phone = c.FormValue("phone")
	userModel.Password = c.FormValue("password")
	userModel.ConfirmPassword = c.FormValue("confirmPassword")

	if err := r.query.Create(&userModel).Error; err != nil {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.FailedMessageSaveTH,
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message:    dto.SuccessMessageSaveTH,
		StatusBool: true,
	})
}

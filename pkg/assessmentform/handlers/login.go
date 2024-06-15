package handlers

import (
	"net/http"

	"github.com/Pugpaprika21/internal/dto"
	"github.com/Pugpaprika21/pkg/assessmentform/models"
	"github.com/Pugpaprika21/pkg/assessmentform/myDTO"
	"github.com/Pugpaprika21/pkg/assessmentform/server"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type login struct {
	query *gorm.DB
}

func NewLogin(server *server.EchoServerEnvironment) *login {
	return &login{
		query: server.Connect.ORM,
	}
}

func (l *login) FormLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"title": "ล็อกอิน",
	})
}

func (l *login) LoginVerify(c echo.Context) error {
	if c.Request().Header.Get("X-Requested-With") != "XMLHttpRequest" {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    "ไม่มีข้อมูลที่ถูกส่งมา",
			StatusBool: false,
		})
	}

	var userModel models.User
	l.query.Where("email = ? AND password = ?", c.FormValue("email"), c.FormValue("password")).First(&userModel)
	if userModel.ID == 0 {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.SuccessMessageDataNotFoundTH,
			StatusBool: false,
		})
	}

	userDTO := myDTO.User{
		ID:              userModel.ID,
		Fullname:        userModel.Fullname,
		Email:           userModel.Email,
		Phone:           userModel.Phone,
		Password:        userModel.Password,
		ConfirmPassword: userModel.ConfirmPassword,
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:    "เข้าสู่ระบบสำเร็จ",
		StatusBool: true,
		Data:       userDTO,
	})
}

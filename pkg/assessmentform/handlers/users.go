package handlers

import (
	"net/http"

	"github.com/Pugpaprika21/pkg/assessmentform/server"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type user struct {
	query *gorm.DB
}

func NewUser(server *server.EchoServerEnvironment) *user {
	return &user{
		query: server.Connect.ORM,
	}
}

func (u *user) Userhome(c echo.Context) error {
	return c.Render(http.StatusOK, "user-home.html", map[string]interface{}{
		"title": "หน้าหลัก",
	})
}

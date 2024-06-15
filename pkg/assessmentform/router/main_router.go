package router

import (
	"github.com/Pugpaprika21/pkg/assessmentform/handlers"
	"github.com/Pugpaprika21/pkg/assessmentform/server"
	"github.com/Pugpaprika21/pkg/assessmentform/tmpl"
	"github.com/labstack/echo/v4"
)

func EchoRouter(e *echo.Echo, server *server.EchoServerEnvironment) {
	e.Renderer = tmpl.NewTemplateRegistry().Load()

	g := e.Group("assessmentform")

	loginhandler := handlers.NewLogin(server)
	g.GET("/login", loginhandler.FormLogin)
	g.POST("/login-verify", loginhandler.LoginVerify)

	registerhandler := handlers.NewRegister(server)
	g.GET("/register", registerhandler.FormRegister)
	g.POST("/register-save", registerhandler.RegisterSave)

	userhandler := handlers.NewUser(server)
	g.GET("/user-home", userhandler.Userhome)

	assessmentshandler := handlers.NewAssessments(server)
	g.GET("/form-save", assessmentshandler.FormSaveAssessments)
	g.POST("/save-data", assessmentshandler.SaveDataAssessments)
	g.GET("/user/:userId/assessments/:assessmentId", assessmentshandler.GetAssessmentsByID)
	g.GET("/user/:userId/assessments", assessmentshandler.GetAllAssessmentsByUserID)
	g.GET("/assessments/:assessmentId", assessmentshandler.FormUpdateAssessments)
	g.PUT("/assessments/:assessmentId", assessmentshandler.UpdateAssessments)
	g.DELETE("/user/:userId/assessments/:assessmentId", assessmentshandler.DeleteAssessmentsByUserID)
}

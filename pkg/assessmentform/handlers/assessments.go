package handlers

import (
	"net/http"
	"net/url"

	"github.com/Pugpaprika21/internal/dto"
	"github.com/Pugpaprika21/internal/utils"
	"github.com/Pugpaprika21/pkg/assessmentform/models"
	"github.com/Pugpaprika21/pkg/assessmentform/myDTO"
	"github.com/Pugpaprika21/pkg/assessmentform/server"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type assessments struct {
	query *gorm.DB
}

func NewAssessments(server *server.EchoServerEnvironment) *assessments {
	return &assessments{
		query: server.Connect.ORM,
	}
}

func (a *assessments) FormSaveAssessments(c echo.Context) error {
	return c.Render(http.StatusOK, "assessment-form.html", map[string]interface{}{
		"title": "ฟอร์มแบบประเมินโรคซึมเศร้า",
	})
}

func (a *assessments) SaveDataAssessments(c echo.Context) error {
	if c.Request().Header.Get("X-Requested-With") != "XMLHttpRequest" {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    "ไม่มีข้อมูลที่ถูกส่งมา",
			StatusBool: false,
		})
	}

	assessmentform := c.FormValue("assessmentform")
	userID := utils.UintFromString(c.FormValue("userId"))

	values, _ := url.ParseQuery(assessmentform)

	var assessments models.Assessments
	assessments.Question1 = utils.IntFromString(values.Get("question_1"))
	assessments.Question2 = utils.IntFromString(values.Get("question_2"))
	assessments.Question3 = utils.IntFromString(values.Get("question_3"))
	assessments.Question4 = utils.IntFromString(values.Get("question_4"))
	assessments.Question5 = utils.IntFromString(values.Get("question_5"))
	assessments.Question6 = utils.IntFromString(values.Get("question_6"))
	assessments.Question7 = utils.IntFromString(values.Get("question_7"))
	assessments.Question8 = utils.IntFromString(values.Get("question_8"))
	assessments.Question9 = utils.IntFromString(values.Get("question_9"))
	assessments.Question10 = utils.IntFromString(values.Get("question_10"))
	assessments.UserID = userID

	if err := a.query.Create(&assessments).Error; err != nil {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.FailedMessageSaveTH,
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:    dto.SuccessMessageSaveTH,
		StatusBool: true,
	})
}

func (a *assessments) GetAssessmentsByID(c echo.Context) error {
	if c.Request().Header.Get("X-Requested-With") != "XMLHttpRequest" {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    "ไม่มีข้อมูลที่ถูกส่งมา",
			StatusBool: false,
		})
	}

	userID := utils.UintFromString(c.Param("userId"))
	assessmentID := utils.UintFromString(c.Param("assessmentId"))

	var assessmentCount int64
	var assessmentsModel models.Assessments
	a.query.Model(&assessmentsModel).Where("user_id = ? AND id = ?", userID, assessmentID).Count(&assessmentCount)
	if assessmentCount == 0 {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.SuccessMessageDataNotFoundTH,
			StatusBool: false,
		})
	}

	assessmentAverageScoreSql := `
		SELECT
			(AVG(question1) + AVG(question2) + AVG(question3) + AVG(question4) +
			AVG(question5) + AVG(question6) + AVG(question7) + AVG(question8) +
			AVG(question9) + AVG(question10)) / 10 AS average_score,
			CASE
				WHEN (AVG(question1) + AVG(question2) + AVG(question3) +
					AVG(question4) + AVG(question5) + AVG(question6) +
					AVG(question7) + AVG(question8) + AVG(question9) +
					AVG(question10)) / 10 > 60 THEN 'โรคซึมเศร้า'
				ELSE 'ไม่เป็นโรคซึมเศร้า'
			END AS depression_status
		FROM assessments
		WHERE user_id = ? AND id = ?;
	`
	var assessmentAverageScore myDTO.AssessmentAverageScoreResultFetchRow
	a.query.Raw(assessmentAverageScoreSql, userID, assessmentID).Scan(&assessmentAverageScore)

	return c.JSON(http.StatusOK, dto.Response{
		Message:    dto.SuccessMessageGetTH,
		StatusBool: true,
		Data:       assessmentAverageScore,
	})
}

func (a *assessments) GetAllAssessmentsByUserID(c echo.Context) error {
	if c.Request().Header.Get("X-Requested-With") != "XMLHttpRequest" {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    "ไม่มีข้อมูลที่ถูกส่งมา",
			StatusBool: false,
		})
	}
	userID := utils.UintFromString(c.Param("userId"))

	var assessmentCount int64
	var assessmentsModel models.Assessments
	if err := a.query.Model(&assessmentsModel).Where("user_id = ?", userID).Count(&assessmentCount).Error; err != nil {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.SuccessMessageDataNotFoundTH,
			StatusBool: false,
		})
	}
	if assessmentCount == 0 {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.SuccessMessageDataNotFoundTH,
			StatusBool: false,
		})
	}

	var assessments []models.Assessments
	if err := a.query.Where("user_id = ?", userID).Order("created_at DESC").Find(&assessments).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:    dto.SuccessMessageDataNotFoundTH,
			StatusBool: false,
		})
	}

	var assessmentAverageScore []myDTO.AssessmentAverageScoreResultFetchRow
	for _, assessment := range assessments {
		averageScore := (assessment.Question1 + assessment.Question2 + assessment.Question3 + assessment.Question4 +
			assessment.Question5 + assessment.Question6 + assessment.Question7 + assessment.Question8 +
			assessment.Question9 + assessment.Question10) / 10

		depressionStatus := "ไม่เป็นโรคซึมเศร้า"
		if averageScore > 60 {
			depressionStatus = "โรคซึมเศร้า"
		}

		assessmentAverageScore = append(assessmentAverageScore, myDTO.AssessmentAverageScoreResultFetchRow{
			ID:               assessment.ID,
			AverageScore:     float32(averageScore),
			DepressionStatus: depressionStatus,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:    dto.SuccessMessageGetTH,
		StatusBool: true,
		Data:       assessmentAverageScore,
	})
}

func (a *assessments) FormUpdateAssessments(c echo.Context) error {
	assessmentID := utils.UintFromString(c.Param("assessmentId"))

	var assessmentsModel models.Assessments
	a.query.Where("id = ?", assessmentID).First(&assessmentsModel)
	return c.Render(http.StatusOK, "assessment-form-update.html", map[string]interface{}{
		"title":        "เเก้ไขฟอร์มแบบประเมินโรคซึมเศร้า",
		"assessmentId": assessmentID,
		"assessments":  assessmentsModel,
	})
}

func (a *assessments) UpdateAssessments(c echo.Context) error {
	if c.Request().Header.Get("X-Requested-With") != "XMLHttpRequest" {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    "ไม่มีข้อมูลที่ถูกส่งมา",
			StatusBool: false,
		})
	}

	assessmentID := utils.UintFromString(c.Param("assessmentId"))
	userID := utils.UintFromString(c.FormValue("userId"))
	assessmentform := c.FormValue("assessmentform")

	values, _ := url.ParseQuery(assessmentform)

	var assessments models.Assessments
	assessments.Question1 = utils.IntFromString(values.Get("question_1"))
	assessments.Question2 = utils.IntFromString(values.Get("question_2"))
	assessments.Question3 = utils.IntFromString(values.Get("question_3"))
	assessments.Question4 = utils.IntFromString(values.Get("question_4"))
	assessments.Question5 = utils.IntFromString(values.Get("question_5"))
	assessments.Question6 = utils.IntFromString(values.Get("question_6"))
	assessments.Question7 = utils.IntFromString(values.Get("question_7"))
	assessments.Question8 = utils.IntFromString(values.Get("question_8"))
	assessments.Question9 = utils.IntFromString(values.Get("question_9"))
	assessments.Question10 = utils.IntFromString(values.Get("question_10"))
	assessments.UserID = userID

	if err := a.query.Where("id = ? AND user_id = ?", assessmentID, userID).Updates(&assessments).Error; err != nil {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.FailedMessageUpdateTH,
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:    dto.SuccessMessageUpdateTH,
		StatusBool: true,
	})
}

func (a *assessments) DeleteAssessmentsByUserID(c echo.Context) error {
	if c.Request().Header.Get("X-Requested-With") != "XMLHttpRequest" {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    "ไม่มีข้อมูลที่ถูกส่งมา",
			StatusBool: false,
		})
	}

	userID := utils.UintFromString(c.Param("userId"))
	assessmentID := utils.UintFromString(c.Param("assessmentId"))

	var assessmentsModel models.Assessments
	if err := a.query.Where("user_id = ? AND id = ?", userID, assessmentID).Delete(&assessmentsModel).Error; err != nil {
		return c.JSON(http.StatusOK, dto.Response{
			Message:    dto.FailedMessageDeleteTH,
			StatusBool: false,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:    dto.SuccessMessageDeleteTH,
		StatusBool: true,
	})
}

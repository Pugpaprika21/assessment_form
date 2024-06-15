package myDTO

type AssessmentAverageScoreResultFetchRow struct {
	ID               uint    `gorm:"column:id" json:"assessmentId"`
	AverageScore     float32 `gorm:"column:average_score" json:"average_score"`
	DepressionStatus string  `gorm:"column:depression_status" json:"depression_status"`
}

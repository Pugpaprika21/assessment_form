package myDTO

type User struct {
	ID              uint   `json:"userId"`
	Fullname        string `json:"fullname"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

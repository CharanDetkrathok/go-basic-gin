package model

type Member struct {
	// Name  string `json:"name" binding:"required" validate:"is_charan"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

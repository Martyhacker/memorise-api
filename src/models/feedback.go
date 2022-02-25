package models

type Feedback struct {
	Id      uint   `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

package ViewModel

import (
	"github.com/goblog/app/models"
)

type UserLogin struct {
	models.BaseModel
	Name    string `json:"name"`
	Account string `json:"account"`
	Age     int    `json:"age"`
	Approve bool   `json:"approve"`
	Phone   string `json:"phone"`
	Line    string `json:"line"`
	Address string `json:"address"`
	City    string `json:"city"`
}
type UserLoginPost struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

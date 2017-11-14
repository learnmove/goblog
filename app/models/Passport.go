package models

type Passport struct {
	BaseModel
	Account string `json:"account" gorm:"not null;unique"`
	Token   string `json:"token" gorm:"not null;unique"`
}

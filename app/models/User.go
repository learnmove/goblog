package models

// . "github.com/goblog/app/models/user"

type User struct {
	BaseModel
	Name     string    `json:"name" gorm:"not null"`
	Account  string    `json:"account" gorm:"not null;unique"`
	Password string    `json:"-" gorm:"not null"`
	Age      int       `json:"age"`
	Articles []Article `json:"articles" `
}

// func (user *User)BeforeSave()error{

// }

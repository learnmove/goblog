package models

type City struct {
	BaseModel
	Name string `json:"name"  gorm:"not null"`
}

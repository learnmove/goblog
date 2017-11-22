package models

type PetCategory struct {
	BaseModel
	Name string `json:"name"  gorm:"not null"`
}

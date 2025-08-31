package models

import "time"

type User struct {
	UserId   uint   `json:"userId" gorm:"primaryKey"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}